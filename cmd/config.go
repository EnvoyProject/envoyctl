// Copyright © 2017 Envoy Project <hello@envoyproject.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
