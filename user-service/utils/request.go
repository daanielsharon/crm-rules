package utils

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	ErrMissingID = &RequestError{
		Message: "Missing user ID",
		Status:  http.StatusBadRequest,
	}
)

type RequestError struct {
	Message string
	Status  int
}

func (e *RequestError) Error() string {
	return e.Message
}

func GetIDFromRequest(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", ErrMissingID
	}
	return id, nil
}

func HandleRequestError(w http.ResponseWriter, err error) {
	if reqErr, ok := err.(*RequestError); ok {
		http.Error(w, reqErr.Message, reqErr.Status)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
