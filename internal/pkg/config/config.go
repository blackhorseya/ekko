package config

import (
	"log"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// NewConfigWithPath serve caller to create a viper.Viper
func NewConfigWithPath(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.SetConfigFile(path)

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	log.Println("Using config file: ", v.ConfigFileUsed())

	return v, nil
}

// NewConfig serve caller to create a viper.Viper
func NewConfig() *viper.Viper {
	return viper.New()
}

// WithPathSet is a provider set for wire
var WithPathSet = wire.NewSet(NewConfigWithPath)

// NewViperSet is a provider set for wire
var NewViperSet = wire.NewSet(NewConfig)
