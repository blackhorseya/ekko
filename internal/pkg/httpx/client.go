package httpx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/google/wire"
)

type client struct {
}

func NewClient() httpx.Client {
	return &client{}
}

func (i *client) Do(req *http.Request) (resp *http.Response, err error) {
	return http.DefaultClient.Do(req)
}

// ClientSet is a provider set for httpx client
var ClientSet = wire.NewSet(NewClient)
