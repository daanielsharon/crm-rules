package scheduler

import (
	"log"
	"strconv"
	"time"

	"worker-service/models"
	"worker-service/publisher"
	"worker-service/storage"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	Storage   *storage.Storage
	Publisher *publisher.Publisher
	Cron      *cron.Cron
	Rules     map[string]cron.EntryID
}

func NewScheduler(storage *storage.Storage, publisher *publisher.Publisher) *Scheduler {
	return &Scheduler{
		Storage:   storage,
		Cron:      cron.New(),
		Rules:     make(map[string]cron.EntryID),
		Publisher: publisher,
	}
}

func (s *Scheduler) Start() {
	go s.scheduleRefresher()

	log.Println("Scheduler started!")
	s.Cron.Start()
}

func (s *Scheduler) scheduleRefresher() {
	for {
		s.refreshRules()
		time.Sleep(1 * time.Minute)
	}
}

func (s *Scheduler) refreshRules() {
	rules, err := s.Storage.GetRules()
	if err != nil {
		log.Printf("Failed to fetch rules: %v", err)
		return
	}

	newRules := s.getNewRules(rules)
	s.scheduleNewRules(newRules)
	s.removeStaleRules(rules)
}

func (s *Scheduler) getNewRules(latestRules []models.Rule) []models.Rule {
	var newRules []models.Rule
	for _, rule := range latestRules {
		ruleIDStr := strconv.Itoa(rule.ID)
		if _, exists := s.Rules[ruleIDStr]; !exists {
			newRules = append(newRules, rule)
		}
	}
	return newRules
}

func (s *Scheduler) scheduleNewRules(newRules []models.Rule) {
	for _, rule := range newRules {
		s.addRuleToScheduler(rule)
	}
}

func (s *Scheduler) addRuleToScheduler(rule models.Rule) {
	cronExpression := mapScheduleToCron(rule.Schedule)
	entryID, err := s.Cron.AddFunc(cronExpression, func() {
		s.processRule(rule)
	})
	if err != nil {
		log.Printf("Failed to schedule rule '%s': %v", rule.Name, err)
		return
	}

	ruleIDStr := strconv.Itoa(rule.ID)
	s.Rules[ruleIDStr] = entryID
	log.Printf("Scheduled new rule: %s", rule.Name)
}

func (s *Scheduler) removeStaleRules(latestRules []models.Rule) {
	latestRuleIDs := make(map[string]struct{})
	for _, rule := range latestRules {
		ruleIDStr := strconv.Itoa(rule.ID)
		latestRuleIDs[ruleIDStr] = struct{}{}
	}

	for ruleID, entryID := range s.Rules {
		if _, exists := latestRuleIDs[ruleID]; !exists {
			s.Cron.Remove(entryID)
			delete(s.Rules, ruleID)
			log.Printf("Removed stale rule: %s", ruleID)
		}
	}
}

func mapScheduleToCron(schedule string) string {
	switch schedule {
	case "every_30_minutes":
		return "*/30 * * * *"
	case "every_minute":
		return "* * * * *"
	case "hourly":
		return "0 * * * *"
	case "daily_at_midnight":
		return "0 0 * * *"
	case "every_5_minutes":
		return "*/5 * * * *"
	case "every_10_minutes":
		return "*/10 * * * *"
	default:
		log.Printf("Unknown schedule: %s. Defaulting to hourly.", schedule)
		return "0 * * * *"
	}
}

func (s *Scheduler) processRule(rule models.Rule) {
	log.Printf("Processing rule: %s", rule.Name)

	for _, action := range rule.Actions {
		err := s.Publisher.PublishTask(publisher.NewTask(
			rule.ID,
			rule.Name,
			rule.Condition,
			action.Action,
		))
		if err != nil {
			log.Printf("Failed to publish task for action %s: %v", action.Action, err)
			continue
		}

		log.Printf("Successfully sent task for action %s to execution service", action.Action)
	}
}
