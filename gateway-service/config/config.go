package config

import (
	"fmt"
	"os"
)

type ServiceURLs struct {
	RulesServiceURL         string
	TaskExecutionServiceURL string
	UserServiceURL          string
}

func NewServiceURLs() *ServiceURLs {
	return &ServiceURLs{
		RulesServiceURL: fmt.Sprintf("http://%s:%s/rules/",
			os.Getenv("RULES_SERVICE_HOST"),
			os.Getenv("RULES_SERVICE_PORT")),
		UserServiceURL: fmt.Sprintf("http://%s:%s/users/",
			os.Getenv("USER_SERVICE_HOST"),
			os.Getenv("USER_SERVICE_PORT")),
	}
}
