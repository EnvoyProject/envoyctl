package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type fileConfig struct {
	EnvoyAddress string `yaml:"envoyaddress"`
	Token        string `yaml:"token"`
}

func readConfigFile() (*fileConfig, error) {
	//fail if it is not set
	if !viper.IsSet("envoyaddress") || !viper.IsSet("token") {
		return nil, fmt.Errorf("config file not present or values not set")
	}
	var co = &fileConfig{
		EnvoyAddress: viper.GetString("envoyaddress"),
		Token:        viper.GetString("token"),
	}
	return co, nil
}
