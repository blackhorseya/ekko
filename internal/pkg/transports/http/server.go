package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/netutils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ServerOptions declare http server configuration
type ServerOptions struct {
	Host string
	Port int
	Mode string
}

// NewServerOptions serve caller to create ServerOptions
func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, nil
}

type server struct {
	host string
	port int

	logger     *zap.Logger
	router     *gin.Engine
	httpserver *http.Server
}

// NewRouter serve caller to create a *gin.Engine instance
func NewRouter(opts *ServerOptions) *gin.Engine {
	// todo: 2022/12/4|sean|handle opts.Mode is invalid
	gin.SetMode(opts.Mode)

	return gin.New()
}

// NewServer return a server instance
func NewServer(opts *ServerOptions, logger *zap.Logger, router *gin.Engine) Server {
	return &server{
		host:       opts.Host,
		port:       opts.Port,
		logger:     logger.With(zap.String("type", "http.server")),
		router:     router,
		httpserver: nil,
	}
}

func (i *server) Start() error {
	if i.host == "" {
		i.host = "0.0.0.0"
	}

	if i.port == 0 {
		i.port = netutils.GetAvailablePort()
	}

	addr := fmt.Sprintf("%s:%d", i.host, i.port)

	i.httpserver = &http.Server{
		Addr:              addr,
		Handler:           i.router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	i.logger.Info("http server starting...", zap.String("addr", addr))

	go func() {
		err := i.httpserver.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			i.logger.Fatal("start http server error", zap.Error(err))
			return
		}
	}()

	return nil
}

func (i *server) Stop() error {
	i.logger.Info("http server stopping")

	timeout, cancel := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancel()

	err := i.httpserver.Shutdown(timeout)
	if err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

// ProviderServerSet is a provider set for http server
var ProviderServerSet = wire.NewSet(NewServerOptions, NewRouter, NewServer)
