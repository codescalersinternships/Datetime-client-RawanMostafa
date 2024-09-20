package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/exp/slog"
)

type Client struct {
	baseUrl  string
	endpoint string
	port     string
	client   http.Client
}

func NewClient(baseUrl string, endpoint string, port string, timeout time.Duration) Client {
	slog.Info("New Client created! \n")
	return Client{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		port:     port,
		client:   http.Client{Timeout: timeout},
	}
}

func (c Client) sendRequest(contentType string) (*http.Response, error) {
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

func (c Client) RetrySendRequest(contentType string) (*http.Response, error) {
	var resp *http.Response
	var err error

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second

	retryError := backoff.RetryNotify(func() error {
		resp, err = c.sendRequest(contentType)
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
