package fakeclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//BaseURLV1 is the base url of the api
const (
	BaseURLV1 = "http://localhost:8080/"
)

//Client is
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

//NewClient creats and returns a fresh client
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type successResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Message string `json:"error_message"`
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) (int, error) {

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/vnd.api+json")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return res.StatusCode, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return res.StatusCode, fmt.Errorf("unknown error, status code:%d", res.StatusCode)
		}

		return res.StatusCode, errors.New(errRes.Message)
	}

	fullResponse := successResponse{
		Data: v,
	}

	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		//If the response body is empty (possible in DELETE request)
		body, _ := ioutil.ReadAll(res.Body)
		if len(string(body)) == 0 {
			return res.StatusCode, nil
		} else {
			return res.StatusCode, err
		}
	}

	return res.StatusCode, nil
}
