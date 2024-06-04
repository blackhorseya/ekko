package linebot

import (
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
	server   *httpx.Server
}

func newRest(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func newService(injector *wirex.Injector, server *httpx.Server) adapterx.Servicer {
	return newRest(injector, server)
}

func (i *impl) Start() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) InitRouting() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) GetRouter() *gin.Engine {
	// TODO implement me
	panic("implement me")
}
