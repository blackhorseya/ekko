package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler is the api handler.
func Handler(g *gin.RouterGroup, authenticator *authx.Authenticator) {
	g.GET("/login", login(authenticator))

	g.GET("/callback", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/me", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/logout", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})
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

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
