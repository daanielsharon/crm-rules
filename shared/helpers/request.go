package helpers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetIDFromRequest(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("user ID is required")
	}
	return id, nil
}
