// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package tokenx

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateTokenizer(opts *Options) Tokenizer {
	tokenizer := NewJwtx(opts)
	return tokenizer
}

// wire.go:

var testProviderSet = wire.NewSet(NewJwtx)