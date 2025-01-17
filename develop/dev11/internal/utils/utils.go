package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"dev11/internal/calendar"
)

type ResultResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// ParseEventParams Parsing and validation of parameters.
func ParseEventParams(r *http.Request) (*calendar.Event, error) {
	if err := r.ParseForm(); err != nil {
		return &calendar.Event{}, errors.New("invalid form data")
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		return &calendar.Event{}, errors.New("id is required")
	}

	title := r.FormValue("title")
	if title == "" {
		return &calendar.Event{}, errors.New("title is required")
	}

	dateStr := r.FormValue("date")
	if dateStr == "" {
		return &calendar.Event{}, errors.New("date is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &calendar.Event{}, errors.New("id must be a valid integer")
	}

	date, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		return &calendar.Event{}, errors.New("date must be in YYYY-MM-DD hh:mm format")
	}

	return &calendar.Event{
		ID:    id,
		Title: title,
		Date:  date,
	}, nil
}

func SendResult(w http.ResponseWriter, response string) error {
	data := ResultResponse{response}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func SendEvents(w http.ResponseWriter, response []calendar.Event) error {
	data := calendar.Result{Result: response}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func SendError(w http.ResponseWriter, err error, statusCode int) {
	data := ErrorResponse{err.Error()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
