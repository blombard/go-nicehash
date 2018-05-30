/*
   client.go
       Wrapper for the Nicehash API
*/
package nicehash

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	API_id     string
	API_key    string
	BTC_addr   string
	httpClient *http.Client
}

type BadRequest struct {
	code int64  `json:"code"`
	msg  string `json:"msg,required"`
}

func handleError(resp *http.Response) error {
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("Bad response Status %s. Response Body: %s", resp.Status, string(body))
	}
	return nil
}

// Creates a new Nicehash HTTP Client
func NewClient(API_id, API_key, BTC_addr string) (c *Client) {
	client := &Client{
		API_id:     API_id,
		API_key:    API_key,
		BTC_addr:   BTC_addr,
		httpClient: &http.Client{},
	}
	return client
}

func (c *Client) do(method, endURL, typeAuth string, result interface{}) (resp *http.Response, err error) {

	switch typeAuth {
	case "private":
		if len(c.API_id) == 0 || len(c.API_key) == 0 {
			err = errors.New("Private endpoints require you to set an API ID and API KEY")
			return
		}
		endURL += "&id=" + c.API_id + "&key=" + c.API_key
	case "addr":
		if len(c.BTC_addr) == 0 {
			err = errors.New("You need to provide a valid wallet address")
			return
		}
		endURL += "&addr=" + c.BTC_addr
	}

	fullUrl := fmt.Sprintf("%s%s", BaseUrl, endURL)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}

	// Check for error
	defer resp.Body.Close()
	err = handleError(resp)
	if err != nil {
		return
	}

	// Process response
	if resp != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(result)
	}
	return
}
