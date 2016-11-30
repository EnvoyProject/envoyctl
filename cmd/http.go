package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//Config stores the configuration of the client
type Config struct {
	//Address points to the address of the blacklist AIP
	Address string
	//HTTPClient is the client to use. Default will be used if not provided
	HTTPClient *http.Client
	//Token is the authentication token to be used for the requests
	Token string
}

// //DefaultConfig returns a default configuration
// func DefaultConfig() *Config {
// 	return defaultConfig()
// }

//Client provides a client for the blacklist API
type Client struct {
	config Config
}

//NewClient creates a new client
func NewClient() (*Client, error) {
	diskconfig, err := readConfigFile()
	if err != nil {
		return nil, err
	}
	config := &Config{
		Address: diskconfig.EnvoyAddress,
		HTTPClient: &http.Client{
			Transport: http.DefaultTransport,
		},
		Token: diskconfig.Token,
	}

	client := &Client{
		config: *config,
	}
	return client, nil
}

// //request is a helper to build a request to the API server
// type request struct {
// 	config *Config
// 	method string
// 	url    *url.URL
// 	params url.Values
// 	body   io.Reader
// 	header http.Header
// 	obj    interface{}
// }

// Request this maps a new request
type Request struct {
	Method string
	Path   string
	Values url.Values
	Body   io.Reader
}

//newRequest is used to make a new request
//and include the configuration data
func (c *Client) newRequest(r *Request) (*http.Request, error) {

	uri := url.URL{
		Scheme: "http",
		Host:   c.config.Address,
		Path:   "/v1.0/" + r.Path,
		//RawQuery: params.Encode
	}
	uri.RawQuery = r.Values.Encode()
	//url.Query you can set parameters
	req, err := http.NewRequest(r.Method, uri.RequestURI(), r.Body)
	if err != nil {
		return nil, err
	}
	req.URL.Scheme = uri.Scheme
	req.URL.Host = uri.Host
	req.Host = uri.Host

	req.Header.Set("X-Envoy-Token", c.config.Token)
	return req, nil
}

//doRequest performs the request
func (c *Client) doRequest(r *http.Request) (time.Duration, *http.Response, error) {
	//output the request if needed
	if debug {
		if err := dumpRequest(r); err != nil {
			log.Fatal(err)
		}
	}
	start := time.Now()
	resp, err := c.config.HTTPClient.Do(r)
	diff := time.Now().Sub(start)
	return diff, resp, err
}

// decodeBody is used to JSON decode a body
func decodeBody(resp *http.Response, out interface{}) error {
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// encodeBody is used to encode a request body
func encodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}

//dumpRequest outputs the request for debugging
func dumpRequest(r *http.Request) error {
	dump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%q", dump)
	return nil
}
