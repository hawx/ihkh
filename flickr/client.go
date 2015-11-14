package flickr

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://api.flickr.com/services/rest/"
	defaultUserAgent = "me.hawx.ihkh"
)

type Client struct {
	client *http.Client
	apiKey string

	BaseURL   *url.URL
	UserAgent string
}

func New(apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		client: http.DefaultClient,
		apiKey: apiKey,

		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}
}

func (client *Client) get(method string, params url.Values, v interface{}) (*http.Response, error) {
	params.Add("method", method)
	params.Add("api_key", client.apiKey)

	reqURL := client.BaseURL.String() + "?" + params.Encode()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", client.UserAgent)

	resp, err := client.client.Do(req)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Received %d response", resp.StatusCode)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if v != nil {
		err = xml.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
