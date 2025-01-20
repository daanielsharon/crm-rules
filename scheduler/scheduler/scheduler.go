package scheduler

import (
	"log"
	"sync"
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
	Rules     map[cron.EntryID]models.Rule
	mu        sync.RWMutex
	stopChan  chan struct{}
}

func NewScheduler(storage *storage.Storage, publisher *publisher.Publisher) *Scheduler {
	return &Scheduler{
		Storage:   storage,
		Cron:      cron.New(),
		Rules:     make(map[cron.EntryID]models.Rule),
		Publisher: publisher,
		stopChan:  make(chan struct{}),
	}
}

func (s *Scheduler) Start() {
	s.Cron.Start()
	go s.periodicRuleSync()

	log.Println("Scheduler started!")
}

func (s *Scheduler) Stop() {
	s.Cron.Stop()
	close(s.stopChan)

	log.Println("Scheduler stopped!")
}

func (s *Scheduler) periodicRuleSync() {
	s.syncRules()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.syncRules()
		case <-s.stopChan:
			return
		}
	}
}

func (s *Scheduler) syncRules() {
	rules, err := s.Storage.GetRules()
	if err != nil {
		log.Printf("Failed to fetch rules: %v", err)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for entryID := range s.Rules {
		s.Cron.Remove(entryID)
	}
	s.Rules = make(map[cron.EntryID]models.Rule)

	for _, rule := range rules {
		s.scheduleRule(rule)
	}

	log.Printf("Synced %d rules", len(rules))
}

func (s *Scheduler) scheduleRule(rule models.Rule) {
	cronExpression := mapScheduleToCron(rule.Schedule)

	entryID, err := s.Cron.AddFunc(cronExpression, func() {
		s.processRule(rule)
	})

	if err != nil {
		log.Printf("Failed to schedule rule '%s': %v", rule.Name, err)
		return
	}

	s.Rules[entryID] = rule
	log.Printf("Scheduled rule: %s", rule.Name)
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
