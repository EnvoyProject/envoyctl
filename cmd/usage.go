// Copyright © 2016 Paul Piscuc <paul.piscuc@gmail.com>
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

// usageCmd represents the usage command
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "View your usage",
	Long: `Access the Envoy API to view your usage
`,
	Run: func(cmd *cobra.Command, args []string) {
		ex := runGetUsage(cmd, args)
		os.Exit(ex)
	},
}

type usageResponse struct {
	Code int64 `json:"code"`
	Data struct {
		MaxCalls int64 `json:"maxcalls"`
		Total    int64 `json:"total"`
	}
}

func runGetUsage(cmd *cobra.Command, args []string) int {
	if len(args) != 0 {
		fmt.Println("usage must be run without parameters")
		return -1
	}
	m, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	req, err := m.newRequest("GET", "usage", nil)
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
	var usageresp usageResponse
	if err = decodeBody(resp, &usageresp); err != nil {
		fmt.Println(err)
		return -1
	}
	switch {
	case output == "json" || output == "":
		printJSON(usageresp.Data)
		return 0
	case output == "csv":
		printCSV(usageresp.Data)
		return 0
	default:
		fmt.Println("output not implemented")
	}
	return 0
}

func init() {
	RootCmd.AddCommand(usageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
