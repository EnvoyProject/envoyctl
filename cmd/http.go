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

//Client stores the configuration of the client
type Client struct {
	//Address points to the address of the searchhead api
	Address string
	//HTTPClient is the client to use. Default will be used if not provided
	HTTPClient *http.Client
	//Token is the authentication token to be used for the requests
	Token string
}

//NewClient creates a new client, and gets a secure token
func NewClient() (*Client, error) {
	diskconfig, err := readConfigFile()
	if err != nil {
		return nil, err
	}
	token, err := getSecureToken(diskconfig.EnvoyAddress, diskconfig.Username, diskconfig.Apikey)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Address: diskconfig.EnvoyAddress,
		HTTPClient: &http.Client{
			Transport: http.DefaultTransport,
		},
		Token: token,
	}

	return client, nil
}

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
		Host:   c.Address,
		Path:   r.Path,
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

	req.Header.Set("X-Envoy-Token", c.Token)
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
	resp, err := c.HTTPClient.Do(r)
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

//getSecureToken gets a temporary token to perform operations
func getSecureToken(backend, username, apikey string) (string, error) {
	apilogin := struct {
		Username string `json:"username"`
		Apikey   string `json:"apikey"`
	}{
		Username: username,
		Apikey:   apikey,
	}
	enc, err := json.Marshal(apilogin)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("http://%s/api/apisignin", backend)
	resp, err := http.Post(url, "application/json", bytes.NewReader(enc))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//handleError(resp)
		return "", fmt.Errorf("response error code: %d", resp.StatusCode)
	}
	type apiResponse struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		ID       int64  `json:"id"`
		Apikey   string `json:"apikey"`
	}
	var data apiResponse
	if err = decodeBody(resp, &data); err != nil {
		return "", err
	}
	return data.Token, nil

}
