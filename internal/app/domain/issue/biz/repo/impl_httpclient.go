package repo

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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
	opts *HTTPClientOptions
}

// NewHTTPClient serve caller to create an IRepo
func NewHTTPClient(opts *HTTPClientOptions) IRepo {
	return &httpclient{
		opts: opts,
	}
}

func (h *httpclient) GetByID(ctx contextx.Contextx, id int64) (info *im.Ticket, err error) {
	// todo: 2023/4/3|sean|impl me
	panic("implement me")
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
