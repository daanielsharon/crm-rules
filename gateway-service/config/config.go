package config

import (
	"fmt"
	"os"
)

type ServiceURLs struct {
	RulesServiceURL         string
	TaskExecutionServiceURL string
}

func NewServiceURLs() *ServiceURLs {
	return &ServiceURLs{
		RulesServiceURL: fmt.Sprintf("http://%s:%s/rules/",
			os.Getenv("RULES_SERVICE_HOST"),
			os.Getenv("RULES_SERVICE_PORT")),
		// TaskExecutionServiceURL: fmt.Sprintf("http://%s:%s/tasks/",
		// 	os.Getenv("TASK_EXECUTION_SERVICE_HOST"),
		// 	os.Getenv("TASK_EXECUTION_SERVICE_PORT")),
	}
}
