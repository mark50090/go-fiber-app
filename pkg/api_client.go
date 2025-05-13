package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL   string
	HTTP      *http.Client
	ApiKey    string
	HeaderKey string
}

func NewClient(baseURL string, apiKey string, headerKey string) *Client {
	return &Client{
		BaseURL:   baseURL,
		HTTP:      &http.Client{},
		ApiKey:    apiKey,
		HeaderKey: headerKey,
	}
}

func (c *Client) Fetch(endpoint string, headers map[string]string, params map[string]string, result interface{}) error {

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add(c.HeaderKey, c.ApiKey)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	q := req.URL.Query()
	for key, value := range params {
		q.Add(strings.ToLower(key), value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(result)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("API call failed: " + resp.Status)
	}

	return nil
}
