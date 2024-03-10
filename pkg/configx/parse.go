package configx

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// C is a global config instance.
var C = new(Config)

// LoadWithPathAndName loads config from path and name.
func LoadWithPathAndName(path, name string) (err error) {
	v := viper.GetViper()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, _ := os.UserHomeDir()
		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName("." + name)
	}

	err = bindEnv(v)
	if err != nil {
		return err
	}

	err = v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}

func bindEnv(v *viper.Viper) (err error) {
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = v.BindEnv("log.level", "LOG_LEVEL")
	if err != nil {
		return err
	}

	err = v.BindEnv("log.format", "LOG_FORMAT")
	if err != nil {
		return err
	}

	err = v.BindEnv("http.mode", "HTTP_MODE")
	if err != nil {
		return err
	}

	err = v.BindEnv("storage.mongodb.dsn", "STORAGE_MONGODB_DSN")
	if err != nil {
		return err
	}

	err = v.BindEnv("linebot.secret", "LINEBOT_SECRET")
	if err != nil {
		return err
	}

	err = v.BindEnv("linebot.token", "LINEBOT_TOKEN")
	if err != nil {
		return err
	}

	return nil
}
