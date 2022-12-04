package http

import (
	"net/http"

	"github.com/google/wire"
)

// NewClient return Client
func NewClient() Client {
	return &http.Client{}
}

// ProviderClientSet is a provider set for http client
var ProviderClientSet = wire.NewSet(NewClient)
