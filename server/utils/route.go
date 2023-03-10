package utils

import "net/http"

// Route defines a single route, e.g. a human-readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Name         string
	Method       string
	Pattern      string
	HandlerFunc  http.HandlerFunc
	Authenticate bool
}
