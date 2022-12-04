package httpx

import (
	"net/http"

	"github.com/blackhorseya/todo-app/pkg/httpx"
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

// ProviderClientSet is a provider set for httpx client
var ProviderClientSet = wire.NewSet(NewClient)
