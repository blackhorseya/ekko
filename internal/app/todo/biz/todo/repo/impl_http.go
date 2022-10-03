package repo

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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

// NewHTTP return ITodoRepo
func NewHTTP(opts *Options, restclient restclient.RestClient) ITodoRepo {
	return &rest{
		opts:       opts,
		restclient: restclient,
	}
}

func (i *rest) GetByID(ctx contextx.Contextx, id uint64) (task *ticket.Task, err error) {
	uri, err := url.Parse(i.opts.BaseURL + "/api/v1/tasks/" + strconv.Itoa(int(id)))
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

	var ret *ticket.Task
	err = json.Unmarshal(str, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *rest) List(ctx contextx.Contextx, condition QueryTodoCondition) (tasks []*ticket.Task, err error) {
	uri, err := url.Parse(fmt.Sprintf("%s/api/v1/tasks?page=%v&size=%v", i.opts.BaseURL, (condition.Limit+1)/10, condition.Limit))
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

	var ret []*ticket.Task
	err = json.Unmarshal(str, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *rest) Create(ctx contextx.Contextx, created *ticket.Task) (task *ticket.Task, err error) {
	uri, err := url.Parse(i.opts.BaseURL + "/api/v1/tasks")
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("title", created.Title)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader([]byte(values.Encode())))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	var ret *ticket.Task
	err = json.Unmarshal(str, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *rest) Count(ctx contextx.Contextx) (total int, err error) {
	uri, err := url.Parse(fmt.Sprintf("%s/api/v1/tasks?page=%v&size=%v", i.opts.BaseURL, 1, 10))
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return 0, err
	}

	resp, err := i.restclient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data *response.Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	if data.Code != 200 {
		return 0, errors.New(strconv.Itoa(data.Code) + " " + data.Msg)
	}

	ret, err := strconv.Atoi(resp.Header.Get("x-total-count"))
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *rest) Update(ctx contextx.Contextx, updated *ticket.Task) (task *ticket.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *rest) Remove(ctx contextx.Contextx, id uint64) error {
	// TODO implement me
	panic("implement me")
}
