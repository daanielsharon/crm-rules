package types

type Task struct {
	RuleID    int    `json:"rule_id"`
	Name      string `json:"name"`
	Condition string `json:"condition"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}
