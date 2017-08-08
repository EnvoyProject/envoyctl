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
	"os"

	"github.com/spf13/cobra"
)

// getindicatorCmd represents the getindicator command
var getindicatorCmd = &cobra.Command{
	Use:   "getindicator",
	Short: "Get the value of an indicator",
	Long:  `Access the Envoy API to get an indicator`,
	Run: func(cmd *cobra.Command, args []string) {
		ex := runGetIndicator(cmd, args)
		os.Exit(ex)
	},
}

func runGetIndicator(cmd *cobra.Command, args []string) int {
	if len(args) < 2 {
		fmt.Println("please provide an indicatortype and an indicator")
		return -1
	}
	if len(args) >= 2 {
		indicatortype := args[0]
		indicator := args[1]
		m, err := NewClient()
		if err != nil {
			fmt.Println(err)
			return -1
		}
		r := &Request{
			Method: "GET",
			Path:   "indicators/" + indicatortype + "/" + indicator,
			Body:   nil,
			//Values
		}
		req, err := m.newRequest(r)
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
		if resp.StatusCode != 200 {
			handleError(resp)
			return 0
		}
		if output == "raw" {
			printRaw(resp.Body)
			return 0
		}
		var indicat arenaIndicator
		if err = decodeBody(resp, &indicat); err != nil {
			fmt.Println(err)
			return -1
		}
		switch {
		case output == "json" || output == "":
			printJSON(indicat)
			return 0
		case output == "csv":
			printCSV(indicat)
			return 0
		default:
			fmt.Println("output not implemented")
		}

		return 0
	}

	return 0
}

func init() {
	RootCmd.AddCommand(getindicatorCmd)
}
