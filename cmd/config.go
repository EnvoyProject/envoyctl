package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type fileConfig struct {
	EnvoyAddress string `yaml:"envoyaddress"`
	Username     string `yaml:"username"`
	Apikey       string `yaml:"apikey"`
}

func readConfigFile() (*fileConfig, error) {
	//fail if it is not set
	if !viper.IsSet("envoyaddress") || !viper.IsSet("apikey") || !viper.IsSet("username") {
		return nil, fmt.Errorf("config file not present or values not set")
	}
	var co = &fileConfig{
		EnvoyAddress: viper.GetString("envoyaddress"),
		Username:     viper.GetString("username"),
		Apikey:       viper.GetString("apikey"),
	}
	return co, nil
}
