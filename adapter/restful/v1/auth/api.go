package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler is the api handler.
func Handler(g *gin.RouterGroup, authenticator *authx.Authenticator) {
	g.GET("/login", login(authenticator))

	g.GET("/callback", callback(authenticator))

	g.GET("/me", me())

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
		ctx, err := contextx.FromGin(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		session := sessions.Default(c)
		if c.Query("state") != session.Get("state") {
			c.JSON(http.StatusBadRequest, response.Err.WithMessage("invalid state"))
			return
		}

		token, err := authenticator.Exchange(ctx, c.Query("code"))
		if err != nil {
			_ = c.Error(err)
			return
		}

		idToken, err := authenticator.VerifyIDToken(ctx, token)
		if err != nil {
			_ = c.Error(err)
			return
		}

		var profile map[string]any
		err = idToken.Claims(&profile)
		if err != nil {
			_ = c.Error(err)
			return
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		err = session.Save()
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "/api/v1/auth/me")
	}
}

func me() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		profile := session.Get("profile")
		if profile == nil {
			c.JSON(http.StatusUnauthorized, response.Err.WithMessage("unauthorized"))
			return
		}

		c.HTML(http.StatusOK, "user.html", profile)
	}
}

func logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutURL, err := url.ParseRequestURI("https://" + configx.A.Auth0.Domain + "/2/logout")
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
		parameters.Add("client_id", configx.A.Auth0.ClientID)
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
