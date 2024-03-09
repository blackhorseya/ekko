package notifier

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type line struct {
	endpoint    string
	accessToken string
}

// NewLineNotifier creates a new Notifier that sends notifications to a LINE group.
func NewLineNotifier() (Notifier, error) {
	_, err := url.ParseRequestURI(configx.C.LineNotify.Endpoint)
	if err != nil {
		return nil, err
	}

	return &line{
		endpoint:    configx.C.LineNotify.Endpoint,
		accessToken: configx.C.LineNotify.AccessToken,
	}, nil
}

func (i *line) SendText(ctx contextx.Contextx, message string) error {
	uri, err := url.ParseRequestURI(i.endpoint)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Add("message", message)
	uri.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+i.accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	var got *response
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return err
	}

	if got.Status != http.StatusOK {
		return errors.New(got.Message)
	}

	return nil
}
