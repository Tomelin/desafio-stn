package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Configs struct {
	Webserver  ConfigWebServer "mapstructure:'webserver'"
	Kubeconfig string          "mapstructure:'kubeconfig'"
	Timeout    time.Duration   "mapstructure:'timeout'"
}

func LoadConfigs() (*Configs, error) {

	path := os.Getenv("PATH_CONFIG")
	if path == "" {
		path = "/app/config"
	}

	var cfg *Configs

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found in %s", path)
		}
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	setEnvPath := path + "config.yaml"
	err = os.Setenv("JSON_CONFIG_PATH", setEnvPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
