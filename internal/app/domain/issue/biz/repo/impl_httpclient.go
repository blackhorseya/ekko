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
	v      *viper.Viper
	client httpx.Client
}

// NewHTTPClient serve caller to create an IRepo
func NewHTTPClient(v *viper.Viper, client httpx.Client) (IRepo, error) {
	return &httpclient{
		v:      v,
		client: client,
	}, nil
}

func (h *httpclient) GetBaseURL() (*url.URL, error) {
	o, err := NewHTTPClientOptions(h.v)
	if err != nil {
		return nil, err
	}

	return url.ParseRequestURI(o.URL)
}

func (h *httpclient) GetByID(ctx contextx.Contextx, id int64) (info *im.Ticket, err error) {
	baseURL, err := h.GetBaseURL()
	if err != nil {
		return nil, err
	}
	uri := baseURL.JoinPath(fmt.Sprintf("/v1/tasks/%v", id))

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
	size := condition.Limit
	page := (condition.Offset / condition.Limit) + 1

	baseURL, err := h.GetBaseURL()
	if err != nil {
		return nil, err
	}
	uri := baseURL.JoinPath("/v1/tasks")
	uri.RawQuery = url.Values{
		"page": []string{fmt.Sprintf("%v", page)},
		"size": []string{fmt.Sprintf("%v", size)},
	}.Encode()

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
		Data struct {
			Total int          `json:"total"`
			List  []*im.Ticket `json:"list"`
		} `json:"data,omitempty"`
	}
	var res *dto
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if len(res.Data.List) == 0 {
		return nil, nil
	}

	return res.Data.List, nil
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
