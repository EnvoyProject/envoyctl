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

// apiversionCmd represents the apiversions command
var apiversionCmd = &cobra.Command{
	Use:   "apiversion",
	Short: "Show version of Envoy API",
	Long:  `Get the version of the Envoy service.`,
	Run: func(cmd *cobra.Command, args []string) {
		ex := runGetAPIVersion(cmd, args)
		os.Exit(ex)
	},
}

type version struct {
	Version string `json:"version"`
}

func runGetAPIVersion(cmd *cobra.Command, args []string) int {
	m, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	r := &Request{
		Method: "GET",
		Path:   "version",
		Body:   nil,
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
	var ver version
	if err = decodeBody(resp, &ver); err != nil {
		fmt.Println(err)
		return -1
	}
	switch {
	case output == "json" || output == "":
		printJSON(ver)
		return 0
	case output == "csv":
		printCSV(ver)
		return 0
	default:
		fmt.Println("output not implemented")
	}
	return 0

}

func init() {
	RootCmd.AddCommand(apiversionCmd)
}
