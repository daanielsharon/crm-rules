package types

type Task struct {
	RuleID    int      `json:"rule_id"`
	Condition string   `json:"condition"`
	Actions   []string `json:"actions"`
}
