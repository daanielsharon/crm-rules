package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail checks if an email address matches a standard format
// This regex supports most common email address patterns, including:
// - Alphanumeric characters before the @ symbol
// - Special characters in the username part (before @): periods, exclamation marks, #, $, %, &, ', *, +, /, =, ?, ^, _, `, {, |, }, ~, and -
// - Subdomain.domain formats (e.g., user@sub.example.com)
// Note: Does not cover all possible email formats defined in RFC 5322
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	return emailRegex.MatchString(email)
}
