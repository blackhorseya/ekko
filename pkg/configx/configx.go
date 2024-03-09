package configx

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/ekko/pkg/logging"
	"github.com/blackhorseya/ekko/pkg/netx"
)

// Config defines the config struct.
type Config struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`

	Log  logging.Config `json:"log" yaml:"log"`
	HTTP HTTP           `json:"http" yaml:"http"`

	Storage struct {
		Mongodb struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"mongodb" yaml:"mongodb"`
	} `json:"storage" yaml:"storage"`

	Finmind struct {
		HTTP  HTTP   `json:"http" yaml:"http"`
		Token string `json:"token" yaml:"token"`
	} `json:"finmind" yaml:"finmind"`

	LineNotify struct {
		Endpoint    string `json:"endpoint" yaml:"endpoint"`
		AccessToken string `json:"access_token" yaml:"accessToken"`
	} `json:"line_notify" yaml:"lineNotify"`

	LineBot struct {
		Secret string `json:"secret" yaml:"secret"`
		Token  string `json:"token" yaml:"token"`
	} `json:"line_bot" yaml:"lineBot"`

	Fugle struct {
		APIKey    string    `json:"api_key" yaml:"apiKey"`
		HTTP      HTTP      `json:"http" yaml:"http"`
		Websocket Websocket `json:"websocket" yaml:"websocket"`
	} `json:"fugle" yaml:"fugle"`
}

func (x *Config) String() string {
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

// Websocket defines the websocket struct.
type Websocket struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}
