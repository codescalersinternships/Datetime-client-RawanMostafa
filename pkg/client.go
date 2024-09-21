// This package implements an http client
package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/exp/slog"
)

// Client holds the configurations of the http client as well as the actual http.Client object
type Client struct {
	baseUrl  string
	endpoint string
	port     string
	client   http.Client
}

// NewClient takes the baseUrl, endpoint, port and timeout and returns a Client object
func NewClient(baseUrl string, endpoint string, port string, timeout time.Duration) Client {
	slog.Info("New Client created! \n")
	return Client{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		port:     port,
		client:   http.Client{Timeout: timeout},
	}
}

// SendRequest takes the wanted content type as a string, creates the request
// and adds the content-type header then send the request and returns the response and an error if exists
func (c Client) SendRequest(contentType string) (*http.Response, error) {
	url := c.baseUrl + ":" + c.port + c.endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("couldn't create the request with url: %s", url)
		return nil, err
	}

	req.Header.Add("content-type", contentType)
	resp, err := c.client.Do(req)

	if err != nil {
		slog.Error("couldn't send the request with url: %s", url)
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	slog.Info("your request has been sent successfully!")

	return resp, nil
}

// RetrySendRequest takes the wanted content type as a string, creates and sends the request and uses the retry mechanism
// for maximum of 10 seconds before the request fails
// it then returns the response and an error if it failed to send for 10 seconds
func (c Client) RetrySendRequest(contentType string) (*http.Response, error) {
	var resp *http.Response
	var err error

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second

	retryError := backoff.RetryNotify(func() error {
		resp, err = c.SendRequest(contentType)
		return err
	}, expBackoff, func(err error, d time.Duration) {
		slog.Warn("Request failed, Retrying ...")
	})

	if retryError != nil {
		slog.Error("failed to make the request after retries: %v", err)
		return resp, fmt.Errorf("failed to make the request after retries: %v", err)
	} else {
		return resp, nil
	}
}
