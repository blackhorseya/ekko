package authx

import (
	"errors"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is a struct that holds the authentication logic.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// NewAuthenticator creates a new Authenticator.
func NewAuthenticator() (*Authenticator, error) {
	ctx := contextx.Background()

	provider, err := oidc.NewProvider(ctx, "https://"+configx.C.Auth0.Domain+"/")
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     configx.C.Auth0.ClientID,
		ClientSecret: configx.C.Auth0.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  configx.C.Auth0.CallbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies the ID token.
func (a *Authenticator) VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	return a.Verifier(&oidc.Config{ClientID: a.ClientID}).Verify(ctx, rawIDToken)
}
