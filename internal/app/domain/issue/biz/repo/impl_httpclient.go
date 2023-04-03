package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/blackhorseya/ekko/pkg/contextx"
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// HTTPClientOptions declare http client options
type HTTPClientOptions struct {
	URL string `json:"url" yaml:"url"`
}

// NewHTTPClientOptions serve caller to create an HTTPClientOptions
func NewHTTPClientOptions(v *viper.Viper) (*HTTPClientOptions, error) {
	var (
		err error
		o   = new(HTTPClientOptions)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	return o, err
}

type httpclient struct {
	opts    *HTTPClientOptions
	baseURL *url.URL
	client  httpx.Client
}

// NewHTTPClient serve caller to create an IRepo
func NewHTTPClient(opts *HTTPClientOptions, client httpx.Client) (IRepo, error) {
	baseURL, err := url.ParseRequestURI(opts.URL)
	if err != nil {
		return nil, errors.Wrap(err, "parse url error")
	}

	return &httpclient{
		opts:    opts,
		baseURL: baseURL,
		client:  client,
	}, nil
}

func (h *httpclient) GetByID(ctx contextx.Contextx, id int64) (info *im.Ticket, err error) {
	uri := h.baseURL.JoinPath(fmt.Sprintf("/v1/tasks/%v", id))

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type dto struct {
		*response.Response
		Data *im.Ticket `json:"data"`
	}
	var res *dto
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		ctx.Error("parse response body error", zap.Error(err))
		return nil, err
	}

	if res.Data == nil {
		return nil, nil
	}

	return res.Data, nil
}

func (h *httpclient) List(ctx contextx.Contextx, condition QueryTicketsCondition) (info []*im.Ticket, err error) {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
}

func (h *httpclient) Create(ctx contextx.Contextx, created *im.Ticket) (info *im.Ticket, err error) {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
}

func (h *httpclient) Count(ctx contextx.Contextx, condition QueryTicketsCondition) (total int, err error) {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
}

func (h *httpclient) Update(ctx contextx.Contextx, updated *im.Ticket) error {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
}

func (h *httpclient) DeleteByID(ctx contextx.Contextx, id int64) error {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
}
