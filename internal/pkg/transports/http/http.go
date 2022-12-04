package http

import (
	"net/http"
)

// Server declare a http server functions
//
//go:generate mockery --all --inpackage
type Server interface {
	// Start a server
	Start() error

	// Stop a server
	Stop() error
}

// Client http client interface
//
//go:generate mockery --all --inpackage
type Client interface {
	// Do ...
	Do(req *http.Request) (resp *http.Response, err error)
}
