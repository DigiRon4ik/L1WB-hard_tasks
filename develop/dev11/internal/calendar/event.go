package calendar

import "time"

// Event represents a scheduled event with an ID, title, and date.
// The fields are serialized to and from JSON format, allowing easy data exchange in web applications.
type Event struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
