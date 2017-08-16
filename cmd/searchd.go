// Copyright © 2017 Paul Piscuc <paul@envoyproject.io>
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

// searchdCmd represents the searchd command
var searchdCmd = &cobra.Command{
	Use:   "searchd",
	Short: "Search descriptors",
	Long:  `Transmit a search pattern to the saerch descriptors api.`,
	Run: func(cmd *cobra.Command, args []string) {
		ex := runSearchDescriptors(cmd, args)
		os.Exit(ex)
	},
}

var qp queryParamsRaw

func init() {
	RootCmd.AddCommand(searchdCmd)
	searchdCmd.PersistentFlags().StringVar(&qp.From, "from", "-30d", "from")
	searchdCmd.PersistentFlags().StringVar(&qp.To, "to", "now", "to")
	searchdCmd.PersistentFlags().StringVar(&qp.Q, "q", "", "query")
	searchdCmd.PersistentFlags().Int64Var(&qp.Offset, "offset", 0, "offset of data")
	searchdCmd.PersistentFlags().BoolVar(&qp.Regex, "regexp", false, "interpret query as regex")
	searchdCmd.PersistentFlags().StringVar(&qp.Format, "format", "rows_json", "format of output")
}

type queryParamsRaw struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Q      string `json:"q,omitempty"`
	Offset int64  `json:"offset,omitempty"`
	Regex  bool   `json:"regex,omitempty"`
	Format string `json:"format,omitempty"`
}

func runSearchDescriptors(cmd *cobra.Command, args []string) int {

	if qp.From == "" || qp.To == "" || qp.Q == "" {
		fmt.Println("please provide to, from and query")
		return -1
	}
	m, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if err != nil {
		fmt.Printf("failed to encode the request: %s", err.Error())
		return -1
	}
	values := new(url.Values)
	
	r := &Request{
		Method: "GET",
		Path:   "/api/events/search",
		Body:   nil,
		Values: url.Values()
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Printf("%s\n", body)
	return 0
}
