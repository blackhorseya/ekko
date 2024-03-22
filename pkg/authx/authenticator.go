package authx

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is a struct that holds the authentication logic.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}
