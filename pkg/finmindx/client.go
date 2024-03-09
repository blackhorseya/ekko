package finmindx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

const (
	dateFormat      = "2006-01-02"
	timeoutDuration = 10 * time.Second
)

type impl struct {
	endpoint string
	token    string
}

// NewClient is used to create a new finmindx client instance.
func NewClient() (Dialer, error) {
	return &impl{
		endpoint: configx.C.Finmind.HTTP.URL,
		token:    configx.C.Finmind.Token,
	}, nil
}

func (i *impl) Do(ctx contextx.Contextx, dataset string, params map[string]string, v any) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	baseURL, err := url.ParseRequestURI(i.endpoint)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("dataset", dataset)
	values.Set("token", i.token)
	for k, v := range params {
		values.Set(k, v)
	}

	baseURL.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(timeout, http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var got Response
	err = json.Unmarshal(data, &got)
	if err != nil {
		return err
	}

	if got.Status != http.StatusOK {
		return errors.New(got.Message)
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) TaiwanStockInfo(ctx contextx.Contextx) (res []*TaiwanStockInfo, err error) {
	type response struct {
		*Response `json:",inline"`
		Data      []*TaiwanStockInfo `json:"data,omitempty"`
	}

	var got *response
	err = i.Do(ctx, "TaiwanStockInfo", nil, &got)
	if err != nil {
		return nil, err
	}

	return got.Data, nil
}

func (i *impl) TaiwanStockPrice(
	ctx contextx.Contextx,
	symbol string,
	start, end time.Time,
) (res []*TaiwanStockPrice, err error) {
	var got *TaiwanStockPriceResponse
	err = i.Do(ctx, "TaiwanStockPrice", map[string]string{
		"data_id":    symbol,
		"start_date": start.Format(dateFormat),
		"end_date":   end.Format(dateFormat),
	}, &got)
	if err != nil {
		return nil, err
	}

	return got.Data, nil
}
