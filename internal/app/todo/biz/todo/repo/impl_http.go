package repo

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Options define options
type Options struct {
	BaseURL string `json:"base_url" yaml:"baseURL"`
}

// NewOptions return *Options
func NewOptions(v *viper.Viper) (*Options, error) {
	ret := new(Options)
	ret.BaseURL = v.GetString("baseURL")

	if ret.BaseURL == "" {
		ret.BaseURL = "https://todo.seancheng.space"
	}

	return ret, nil
}

type rest struct {
	opts       *Options
	restclient restclient.RestClient
}

// NewHTTP return IRepo
func NewHTTP(opts *Options, restclient restclient.RestClient) IRepo {
	return &rest{
		opts:       opts,
		restclient: restclient,
	}
}

func (i *rest) GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error) {
	uri, err := url.Parse(i.opts.BaseURL + "/api/v1/tasks/" + id.Hex())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.restclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *response.Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Code != 200 {
		return nil, errors.New(strconv.Itoa(data.Code) + " " + data.Msg)
	}

	str, err := json.Marshal(data.Data)
	if err != nil {
		return nil, err
	}

	var ret *todo.Task
	err = json.Unmarshal(str, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *rest) List(ctx contextx.Contextx, limit, offset int) (tasks []*todo.Task, err error) {
	uri, err := url.Parse(fmt.Sprintf("%s/api/v1/tasks?page=%v&size=%v", i.opts.BaseURL, (offset+1)/10, limit))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.restclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *response.Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Code != 200 {
		return nil, errors.New(strconv.Itoa(data.Code) + " " + data.Msg)
	}

	str, err := json.Marshal(data.Data)
	if err != nil {
		return nil, err
	}

	var ret []*todo.Task
	err = json.Unmarshal(str, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *rest) Create(ctx contextx.Contextx, newTask *todo.Task) (task *todo.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *rest) Count(ctx contextx.Contextx) (total int, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *rest) Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *rest) Remove(ctx contextx.Contextx, id primitive.ObjectID) error {
	// TODO implement me
	panic("implement me")
}
