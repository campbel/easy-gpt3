package openai

import "net/http"

type Client struct {
	http  *http.Client
	token string
}

const (
	apiURL = "https://api.openai.com/v1"
)

func NewClient(http *http.Client, token string) *Client {
	return &Client{http, token}
}

func getURL(endpoint string) string {
	return apiURL + endpoint
}

func or[T comparable](val T, def T) T {
	var zero T
	if val == zero {
		return def
	}
	return val
}
