// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// getindicatorsCmd represents the getindicators command
var getindicatorsCmd = &cobra.Command{
	Use:   "getindicators",
	Short: "Get a list of indicators",
	Long:  `Access the Envoy API to get a list of indicators.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		ex := runGetIndicators(cmd, args)
		os.Exit(ex)
	},
}

func runGetIndicators(cmd *cobra.Command, args []string) int {
	if len(args) != 1 {
		fmt.Println("please provide an indicatortype")
		return -1
	}
	indicatortype := args[0]
	m, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	req, err := m.newRequest("GET", "indicators/"+indicatortype, nil)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	_, resp, err := m.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer resp.Body.Close()

	dataraw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	fmt.Printf("%s\n", dataraw)
	// var data []interface{}
	// if err = decodeBody(resp, data); err != nil {
	// 	c.UI.Error(err.Error())
	// 	return -1
	// }
	//	fmt.Println(data)
	return 0

}

func init() {
	RootCmd.AddCommand(getindicatorsCmd)
}
