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
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/structs"
)

//jsonPretty outputs in an indented json format
func jsonPretty(v interface{}) ([]byte, error) {
	out, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return nil, err
	}
	return out, nil
}

//csvOutput outputs the struct in a csv format
func csvOutput(v interface{}) ([]byte, error) {
	//we need to get the struct into [][]string

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	headers := snames(v)
	values := svalues(v)

	if err := w.Write(headers); err != nil {
		return nil, err
	}
	if err := w.Write(values); err != nil {
		return nil, err
	}
	w.Flush()
	return b.Bytes(), nil
}

//snames gets the names in a struct
func snames(v interface{}) []string {
	names := structs.Names(v)
	return names
}

//svalues gets the values of a struct into a []string
func svalues(v interface{}) []string {
	var out []string
	stru := structs.New(v)
	if stru.IsZero() {
		return out
	}
	for _, val := range stru.Fields() {
		v := val.Value()
		switch v.(type) {
		case string:
			out = append(out, v.(string))
		case int64:
			out = append(out, fmt.Sprintf("%d", (v.(int64))))
		case float64:
			out = append(out, fmt.Sprintf("%.2f", (v.(float64))))
		case time.Time:
			out = append(out, v.(time.Time).Format("2006-01-02 15:04:05"))
		}
	}
	return out
}

//printRaw outputs the response as it is
func printRaw(body io.Reader) {
	dataraw, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", dataraw)

}

//printJSON prints the data into a JSON format
func printJSON(v interface{}) {
	out, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", out)
}

//printCEF prints the data in CEF format
func printCEF(v interface{}) {
	fmt.Printf("cef:%v", v)
}

//printCSV prints the data as csv
func printCSV(v interface{}) {
	out, err := csvOutput(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", out)
}

type errorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

//handleError prints an error based on output formats
func handleError(resp *http.Response) {
	if output == "raw" {
		printRaw(resp.Body)
		return
	}
	var errresp errorResponse
	if err := decodeBody(resp, &errresp); err != nil {
		fmt.Println(err)
		return
	}
	switch {
	case output == "json" || output == "":
		printJSON(errresp)
	case output == "csv":
		printCSV(errresp)
	case output == "cef":
		printCEF(errresp)
	default:
		fmt.Println("error output format not defined")
	}
	return

}
