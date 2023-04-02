//go:generate mockgen -destination=./mock_httpx.go -package=httpx -source=httpx.go

package httpx

import (
	"net/http"
)

// Client http client interface
type Client interface {
	// Do send an HTTP request and returns an HTTP response, following
	// policy (such as redirects, cookies, auth) as configured on the
	// client.
	Do(req *http.Request) (resp *http.Response, err error)
}

// Server declare a http server functions
type Server interface {
	// Start a server
	Start() error

	// Stop a server
	Stop() error
}
