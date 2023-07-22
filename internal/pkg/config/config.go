package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var defaultConfig = &Config{
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
		DSN:   "root:changeme@tcp(localhost:3306)/ekko?charset=utf8mb4&parseTime=True&loc=Local",
		Conns: 100,
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
	DSN   string `json:"dsn" yaml:"dsn"`
	Conns int    `json:"conns" yaml:"conns"`
}

// Jwt is a struct
type Jwt struct {
	Issuer    string `json:"issuer" yaml:"issuer"`
	Signature string `json:"signature" yaml:"signature"`
}

// Config is a struct
type Config struct {
	Log  Log  `json:"log" yaml:"log"`
	HTTP HTTP `json:"http" yaml:"http"`
	DB   DB   `json:"db" yaml:"db"`
	Jwt  Jwt  `json:"jwt" yaml:"jwt"`
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
