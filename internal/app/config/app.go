package config

import "github.com/spf13/viper"

// Config configure application parameters
type Config struct {
	RunMode string
	Swagger Swagger
	HTTP    HTTP
}

// NewConfig is a constructor of config with config path
func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// HTTP configure HTTP server parameters
type HTTP struct {
	Host string
	Port int
}

// Swagger configure swagger parameters
type Swagger struct {
	Enabled bool
}

// Log configure log parameters
type Log struct {
	Level string
}
