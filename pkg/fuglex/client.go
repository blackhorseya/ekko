package fuglex

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/pkg/errors"
)

type dialerResponse struct {
	Code    int    `json:"statusCode,omitempty"`
	Message string `json:"message,omitempty"`
}

// Client is a struct that represents the client.
type Client struct {
	endpoint string
}

// NewClient is a function that is used to create a new client.
func NewClient() (Dialer, error) {
	ep, err := url.ParseRequestURI(configx.C.Fugle.HTTP.URL)
	if err != nil {
		return nil, errors.Wrap(err, "parse fugle http url failed")
	}

	return &Client{
		endpoint: ep.String(),
	}, nil
}

func (c *Client) IntradayQuote(ctx contextx.Contextx, symbol string) (res *IntradayQuote, err error) {
	ep, err := url.ParseRequestURI(c.endpoint + "/intraday/quote/" + symbol)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", configx.C.Fugle.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var got dialerResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		if err != nil {
			return nil, errors.Wrap(err, "decode response body to dialerResponse failed")
		}

		return nil, errors.Errorf("get response failed, code: %d, message: %s", got.Code, got.Message)
	}

	var got *IntradayQuote
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, errors.Wrap(err, "decode response body to IntradayQuote failed")
	}

	return got, nil
}
