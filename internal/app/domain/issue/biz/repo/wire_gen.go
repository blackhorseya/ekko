// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// Injectors from wire.go:

func CreateMariadb(rw *sqlx.DB) IRepo {
	iRepo := NewMariadb(rw)
	return iRepo
}

func CreateHTTPClient(opts *HTTPClientOptions) IRepo {
	iRepo := NewHTTPClient(opts)
	return iRepo
}

// wire.go:

var testProviderMariadbSet = wire.NewSet(NewMariadb)

var testProviderHTTPClientSet = wire.NewSet(NewHTTPClient)