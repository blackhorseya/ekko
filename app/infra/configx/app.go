package configx

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/ekko/pkg/logging"
	"github.com/blackhorseya/ekko/pkg/netx"
	"github.com/google/uuid"
)

// Application defines the application struct.
type Application struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`

	Log  logging.Config `json:"log" yaml:"log"`
	HTTP HTTP           `json:"http" yaml:"http"`

	Storage struct {
		Mongodb struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"mongodb" yaml:"mongodb"`
	} `json:"storage" yaml:"storage"`

	LineNotify struct {
		Endpoint    string `json:"endpoint" yaml:"endpoint"`
		AccessToken string `json:"access_token" yaml:"accessToken"`
	} `json:"line_notify" yaml:"lineNotify"`

	LineBot struct {
		Secret string `json:"secret" yaml:"secret"`
		Token  string `json:"token" yaml:"token"`
	} `json:"line_bot" yaml:"lineBot"`

	Auth0 struct {
		Enabled      bool     `json:"enabled" yaml:"enabled"`
		Domain       string   `json:"domain" yaml:"domain"`
		ClientID     string   `json:"client_id" yaml:"clientID"`
		ClientSecret string   `json:"client_secret" yaml:"clientSecret"`
		CallbackURL  string   `json:"callback_url" yaml:"callbackURL"`
		Audiences    []string `json:"audiences" yaml:"audiences"`
	} `json:"auth0" yaml:"auth0"`

	OTel struct {
		Target string `json:"target" yaml:"target"`
	} `json:"otel" yaml:"otel"`
}

// GetID is used to get the application id.
func (x *Application) GetID() string {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}

	return x.ID
}

func (x *Application) String() string {
	bytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// HTTP defines the http struct.
type HTTP struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr is used to get the http address.
func (http *HTTP) GetAddr() string {
	if http.Host == "" {
		http.Host = "0.0.0.0"
	}

	if http.Port == 0 {
		http.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}
