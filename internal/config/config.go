package config

import (
	"github.com/huobazi/hookme/pkg/voiceover"
	"github.com/spf13/viper"
)

var (
	Config config
)

type config struct {
	Server serverConfig
	Hooks  map[string]hookTask
}

type serverConfig struct {
	Host string
	Port int
}

type hookTask struct {
	Name        string
	Type        string
	RequestPath string
	WorkDir     string
	Command     string
}

func init() {
	viper.SetConfigName("hookme")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.hookme")
	viper.AddConfigPath("/etc/hookme/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			voiceover.Sayf("Not found hookme config file.\n%s \n", err)
			return
		} else {
			voiceover.Sayf("Error reading hookme config file.\n %s", err)
			return
		}
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		voiceover.Sayf("Unable to decode the hookme config file.\n %v", err)
	}
}
