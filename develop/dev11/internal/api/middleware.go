package api

import (
	"log"
	"net/http"
	"time"
)

// Middleware represents an HTTP middleware that wraps around a handler
// to provide additional functionality, such as logging the request method.
type Middleware struct {
	next http.Handler // The next handler in the chain to be executed.
}

// ServeHTTP implements the http.Handler interface, allowing Middleware
// to intercept HTTP requests, log relevant information, and pass the
// request to the next handler in the chain.
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	m.next.ServeHTTP(w, r)
	log.Printf("%s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
}

// NewMiddleware creates a new instance of Middleware with the specified next handler.
func NewMiddleware(next http.Handler) *Middleware {
	return &Middleware{next: next}
}
