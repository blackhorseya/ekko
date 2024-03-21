package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server is an HTTP server.
type Server struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewServer is used to create a new HTTP server.
func NewServer() (*Server, error) {
	ctx := contextx.Background()

	gin.SetMode(configx.C.HTTP.Mode)

	router := gin.New()
	router.Use(ginzap.GinzapWithConfig(ctx.Logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz"},
		Context:    nil,
	}))
	router.Use(ginzap.CustomRecoveryWithZap(ctx.Logger, true, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Err.WithMessage(fmt.Sprintf("%v", err)))
	}))
	router.Use(contextx.AddContextxMiddleware())
	router.Use(response.AddErrorHandlingMiddleware())

	httpserver := &http.Server{
		Addr:              configx.C.HTTP.GetAddr(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}

	return &Server{
		httpserver: httpserver,
		Router:     router,
	}, nil
}

// Start is used to start the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	go func() {
		err := s.httpserver.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("start http server error", zap.Error(err))
		}
	}()

	return nil
}

// Stop is used to stop the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	err := s.httpserver.Shutdown(timeout)
	if err != nil {
		return err
	}

	return nil
}
