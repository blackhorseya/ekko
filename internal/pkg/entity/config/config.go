package config

import (
	"encoding/json"
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Config declare configuration for application
type Config struct {
	APP  *APP  `json:"app" yaml:"app"`
	HTTP *HTTP `json:"http" yaml:"http"`
	DB   *DB   `json:"db" yaml:"db"`
	Log  *Log  `json:"log" yaml:"log"`
}

func (c *Config) String() string {
	ret, _ := json.MarshalIndent(c, "", "  ")
	return string(ret)
}

// APP declare information of application
type APP struct {
	Name       string `json:"name" yaml:"name"`
	Signature string `json:"signature" yaml:"signature"`
}

// HTTP declare http configuration
type HTTP struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddress serve caller to get combine host and port, format is `host:port`
func (h *HTTP) GetAddress() string {
	return fmt.Sprintf("%v:%v", h.Host, h.Port)
}

// DB declare database configuration
type DB struct {
	URL   string `json:"url" yaml:"url"`
	Debug bool   `json:"debug" yaml:"debug"`
}

// Log declare log configuration
type Log struct {
	Level string `json:"level" yaml:"level"`
}

// New serve caller to create a viper.Viper
func New(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New)
