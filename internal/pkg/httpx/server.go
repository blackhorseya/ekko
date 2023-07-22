package httpx

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewRouter() *gin.Engine {
	return gin.New()
}

type server struct {
	host string
	port int

	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewServer(config *config.Config, logger *zap.Logger, router *gin.Engine) httpx.Server {
	gin.SetMode(config.HTTP.Mode)

	return &server{
		host:       config.HTTP.Host,
		port:       config.HTTP.Port,
		logger:     logger,
		router:     router,
		httpServer: http.Server{},
	}
}

func (s *server) Start() error {
	if s.host == "" {
		s.host = "0.0.0.0"
	}

	if s.port == 0 {
		s.port = GetAvailablePort()
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	s.logger.Info("http server starting...")

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal("start http server error", zap.Error(err))
			return
		}
	}()

	s.logger.Info("http server started", zap.String("addr", addr))

	return nil
}

func (s *server) Stop() error {
	s.logger.Info("http server stopping...")

	timeout, cancelFunc := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancelFunc()

	err := s.httpServer.Shutdown(timeout)
	if err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	s.logger.Info("http server stopped")

	return nil
}

// GetAvailablePort returns a port at random
func GetAvailablePort() int {
	l, _ := net.Listen("tcp", ":0") // listen on localhost
	defer l.Close()
	port := l.Addr().(*net.TCPAddr).Port

	return port
}

// ServerSet is a provider set for httpx server
var ServerSet = wire.NewSet(NewRouter, NewServer)
