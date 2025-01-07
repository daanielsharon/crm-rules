package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

func RateLimiter() func(next http.Handler) http.Handler {
	// 100 requests per minute per IP
	return httprate.LimitByIP(100, 1*time.Minute)
}
