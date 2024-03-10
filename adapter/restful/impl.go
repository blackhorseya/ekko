package restful

import (
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
)

type impl struct {
	server *httpx.Server
}

func newRestful(server *httpx.Server) adapterx.Restful {
	return &impl{server: server}
}

func newService(server *httpx.Server) adapterx.Servicer {
	return newRestful(server)
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
