package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var defaultConfig = &Config{
	App: Application{
		Name:     "ekko",
		Version:  "0.0.1",
		Endpoint: "http://localhost:1992",
	},
	Log: Log{
		Level:  "info",
		Format: "json",
	},
	HTTP: HTTP{
		Host: "0.0.0.0",
		Port: 1992,
		Mode: "debug",
	},
	DB: DB{
		DSN:    "root:changeme@tcp(localhost:3306)/ekko?charset=utf8mb4&parseTime=True&loc=Local",
		Conns:  100,
		Source: "github://blackhorseya/ekko/scripts/migrations",
	},
	Jwt: Jwt{
		Issuer:    "ekko",
		Signature: "changeme",
	},
}

// Log is a struct
type Log struct {
	Level  string `json:"level" yaml:"level"`
	Format string `json:"format" yaml:"format"`
}

// HTTP is a struct
type HTTP struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// DB is a struct
type DB struct {
	DSN    string `json:"dsn" yaml:"dsn"`
	Conns  int    `json:"conns" yaml:"conns"`
	Source string `json:"source" yaml:"source"`
}

// Jwt is a struct
type Jwt struct {
	Issuer    string `json:"issuer" yaml:"issuer"`
	Signature string `json:"signature" yaml:"signature"`
}

// Application is a struct
type Application struct {
	Name     string `json:"name" yaml:"name"`
	Version  string `json:"version" yaml:"version"`
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

// Config is a struct
type Config struct {
	App  Application `json:"app" yaml:"app"`
	Log  Log         `json:"log" yaml:"log"`
	HTTP HTTP        `json:"http" yaml:"http"`
	DB   DB          `json:"db" yaml:"db"`
	Jwt  Jwt         `json:"jwt" yaml:"jwt"`
}

// NewWithPath creates a new config with path
func NewWithPath(path string) (*Config, error) {
	if path == "" {
		return defaultConfig, nil
	}

	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigFile(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "read config file failed")
	}

	config := defaultConfig
	err = v.Unmarshal(config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal config failed")
	}

	return config, nil
}

// NewConfigWithViper creates a new config with viper
func NewConfigWithViper(v *viper.Viper) (*Config, error) {
	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "read config failed")
	}

	config := defaultConfig
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal config failed")
	}

	return config, nil
}
