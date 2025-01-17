package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"dev11/internal/utils"
)

// createEventHandler handles the creation of a new event by processing incoming HTTP POST requests.
// It validates the request method and Content-Type, parses the event parameters,
// invokes the business logic to create the event, and sends an appropriate response back to the client.
func (s *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodPost, "application/x-www-form-urlencoded") {
		return
	}

	// Parsing and creating an event object.
	event, err := utils.ParseEventParams(r)
	if err != nil {
		log.Println("Error parsing form:", err)
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	// Calling business logic.
	s.calendar.CreateEvent(event)

	// Return a successful response.
	if err = utils.SendResult(w, "event created successfully"); err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// updateEventHandler processes HTTP POST requests to update an existing event.
// It validates the request method and Content-Type, parses the event parameters from the form data,
// and calls the calendar storage to update the event, sending appropriate responses based on the outcome.
func (s *Server) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodPost, "application/x-www-form-urlencoded") {
		return
	}

	// Parsing and creating an event object.
	event, err := utils.ParseEventParams(r)
	if err != nil {
		log.Println("Error parsing form:", err)
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	// Calling business logic.
	err = s.calendar.UpdateEvent(event.ID, event.Title, event.Date)
	if err != nil {
		log.Println("Error updating data:", err)
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	// Return a successful response.
	if err = utils.SendResult(w, "event updated successfully"); err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// deleteEventHandler handles the deletion of an event based on the provided ID in an HTTP POST request.
// It validates the request method and Content-Type, parses the form data to extract the event ID,
// and calls the calendar storage to delete the event, sending appropriate responses based on the outcome.
func (s *Server) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodPost, "application/x-www-form-urlencoded") {
		return
	}

	// Parse and get ID.
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		utils.SendError(w, errors.New("invalid form data"), http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("Error parsing form:", err)
		utils.SendError(w, errors.New("id must be a valid integer"), http.StatusBadRequest)
		return
	}

	// Calling business logic.
	deleted, err := s.calendar.DeleteEvent(ID)
	if err != nil {
		log.Println("Error deleting data:", err)
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	// Return a successful response.
	err = utils.SendResult(w, fmt.Sprintf("event №%d [%s, %v] deleted", deleted.ID, deleted.Title, deleted.Date))
	if err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// getDailyEventHandler handles HTTP GET requests to retrieve events scheduled for the current day.
func (s *Server) getDailyEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodGet, "") {
		return
	}

	// Calling business logic.
	events := s.calendar.DailyEvents()

	// Return a successful response.
	err := utils.SendEvents(w, events)
	if err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// getWeeklyEventHandler handles HTTP GET requests to retrieve events scheduled for the current week.
func (s *Server) getWeeklyEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodGet, "") {
		return
	}

	// Calling business logic.
	events := s.calendar.WeeklyEvents()

	// Return a successful response.
	err := utils.SendEvents(w, events)
	if err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// getMonthlyEventHandler handles HTTP GET requests to get events scheduled for the current month.
func (s *Server) getMonthlyEventHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the request method and Content-Type.
	if !validateRequest(w, r, http.MethodGet, "") {
		return
	}

	// Calling business logic.
	events := s.calendar.MonthlyEvents()

	// Return a successful response.
	err := utils.SendEvents(w, events)
	if err != nil {
		log.Println("Error writing response:", err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
}

// validateRequest validates the HTTP method and Content-Type of a request.
func validateRequest(w http.ResponseWriter, r *http.Request, expectedMethod, expectedContentType string) bool {
	// Checking the request method.
	if r.Method != expectedMethod {
		log.Println("Invalid HTTP method used:", r.Method)
		errMessage := fmt.Sprintf("only %s method allowed", expectedMethod)
		utils.SendError(w, errors.New(errMessage), http.StatusMethodNotAllowed)
		return false
	}

	// Checking Content-Type.
	if r.Header.Get("Content-Type") != expectedContentType && expectedContentType != "" {
		log.Println("Invalid Content-Type:", r.Header.Get("Content-Type"))
		errMessage := fmt.Sprintf("invalid Content-Type, expected %s", expectedContentType)
		utils.SendError(w, errors.New(errMessage), http.StatusBadRequest)
		return false
	}

	return true
}
