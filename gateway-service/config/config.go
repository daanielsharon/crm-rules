package config

import (
	"fmt"
	"os"
)

type ServiceURLs struct {
	RulesServiceURL string
	UserServiceURL  string
	LogServiceURL   string
}

func NewServiceURLs() *ServiceURLs {
	return &ServiceURLs{
		RulesServiceURL: fmt.Sprintf("http://%s:%s/rules/",
			os.Getenv("RULES_SERVICE_HOST"),
			os.Getenv("RULES_SERVICE_PORT")),
		UserServiceURL: fmt.Sprintf("http://%s:%s/users/",
			os.Getenv("USER_SERVICE_HOST"),
			os.Getenv("USER_SERVICE_PORT")),
		LogServiceURL: fmt.Sprintf("http://%s:%s/logs/",
			os.Getenv("LOG_SERVICE_HOST"),
			os.Getenv("LOG_SERVICE_PORT")),
	}
}
