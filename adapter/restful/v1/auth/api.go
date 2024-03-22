package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler is the api handler.
func Handler(g *gin.RouterGroup, authenticator *authx.Authenticator) {
	g.GET("/login", login(authenticator))

	g.GET("/callback", callback(authenticator))

	g.GET("/me", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/logout", logout())
}

func login(authenticator *authx.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			_ = c.Error(err)
			return
		}

		session := sessions.Default(c)
		session.Set("state", state)
		err = session.Save()
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, authenticator.AuthCodeURL(state))
	}
}

func callback(authenticator *authx.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	}
}

func logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutURL, err := url.ParseRequestURI("https://" + configx.C.Auth0.Domain + "/2/logout")
		if err != nil {
			_ = c.Error(err)
			return
		}

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.ParseRequestURI(scheme + "://" + c.Request.Host)
		if err != nil {
			_ = c.Error(err)
			return
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", configx.C.Auth0.ClientID)
		logoutURL.RawQuery = parameters.Encode()

		c.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
