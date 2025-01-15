package types

type Log struct {
	RuleID     int    `json:"rule_id"`
	UserID     string `json:"user_id"`
	Action     string `json:"action"`
	Status     string `json:"status"`
	ExecutedAt string `json:"executed_at"`
}

func NewLog(ruleID int, userID, action, status string) Log {
	return Log{
		RuleID: ruleID,
		UserID: userID,
		Action: action,
		Status: status,
	}
}
