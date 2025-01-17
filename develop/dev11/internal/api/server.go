package api

import (
	"log"
	"net/http"
	"time"

	"dev11/internal/calendar"
)

// Storage defines an interface for managing calendar events,
// providing methods to create, update, delete, and retrieve events on a daily, weekly, or monthly basis.
type Storage interface {
	CreateEvent(event *calendar.Event)
	UpdateEvent(ID int, Title string, Date time.Time) error
	DeleteEvent(ID int) (*calendar.Event, error)
	DailyEvents() []calendar.Event
	WeeklyEvents() []calendar.Event
	MonthlyEvents() []calendar.Event
}

// Server represents the main application server, encapsulating
// the configuration, HTTP router, middleware, and storage for calendar events.
// It is responsible for handling incoming requests and managing the application’s lifecycle.
type Server struct {
	config     *Config
	router     *http.ServeMux
	middleware *Middleware
	calendar   Storage
}

// NewServer initializes a new Server instance with the provided configuration,
// setting up the HTTP router and middleware, and creating a new calendar storage instance.
func NewServer(config *Config) *Server {
	router := http.NewServeMux()

	return &Server{
		config:     config,
		router:     router,
		middleware: NewMiddleware(router),
		calendar:   calendar.NewCalendar(),
	}
}

// Start begins listening for incoming HTTP requests on the configured address and port,
// logging the server's start message and configuring the router for handling requests.
func (s *Server) Start() error {
	log.Println("Starting API Server on port", s.config.AddrPort)
	s.configureRouter()
	return http.ListenAndServe(s.config.AddrPort, s.middleware)
}

// configureRouter sets up the HTTP routes for the Server,
// mapping specific URL paths and HTTP methods to their corresponding handler functions.
func (s *Server) configureRouter() {
	s.router.HandleFunc("POST /create_event", s.createEventHandler)
	s.router.HandleFunc("POST /update_event", s.updateEventHandler)
	s.router.HandleFunc("POST /delete_event", s.deleteEventHandler)

	s.router.HandleFunc("GET /events_for_day", s.getDailyEventHandler)
	s.router.HandleFunc("GET /events_for_week", s.getWeeklyEventHandler)
	s.router.HandleFunc("GET /events_for_month", s.getMonthlyEventHandler)
}
