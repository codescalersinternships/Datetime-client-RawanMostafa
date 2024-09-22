// This package implements an http client
package pkg

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/exp/slog"
)

// Client holds the configurations of the http client as well as the actual http.Client object
type Client struct {
	baseUrl     string
	endpoint    string
	port        string
	contentType string
	client      http.Client
}

// NewClient takes the baseUrl, endpoint, port, content-type and timeout and returns a Client object
func NewClient(baseUrl string, endpoint string, port string, contentType string, timeout time.Duration) Client {
	slog.Info("New Client created! \n")
	return Client{
		baseUrl:     baseUrl,
		endpoint:    endpoint,
		port:        port,
		contentType: contentType,
		client:      http.Client{Timeout: timeout},
	}
}

func (c Client) getTime() (*http.Response, error) {
	url := c.baseUrl + ":" + c.port + c.endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("couldn't create the request with url: %s", url)
		return nil, err
	}

	req.Header.Add("content-type", c.contentType)
	resp, err := c.client.Do(req)

	if err != nil {
		slog.Error("couldn't send the request with url: %s", url)
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	slog.Info("your request has been sent successfully!")

	return resp, nil
}

func readBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading request body: %v", err)

	}
	return string(body), nil
}

// GetTime creates and sends the request and uses the retry mechanism
// for maximum of 10 seconds before the request fails
// it then returns the time and an error if it failed to send for 10 seconds
func (c Client) GetTime() (time.Time, error) {
	var resp *http.Response
	var err error

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second

	retryError := backoff.RetryNotify(func() error {
		resp, err = c.getTime()
		return err
	}, expBackoff, func(err error, d time.Duration) {
		slog.Warn("Request failed, Retrying ...")
	})

	if retryError != nil {
		slog.Error("failed to make the request after retries: %v", err)
		return time.Time{}, fmt.Errorf("failed to make the request after retries: %v", err)
	}
	body, err := readBody(resp)
	if err != nil {
		return time.Time{}, err
	}
	if resp.StatusCode == http.StatusUnsupportedMediaType {
		return time.Time{}, fmt.Errorf("%s", http.StatusText(http.StatusUnsupportedMediaType))
	}
	body = strings.Trim(body, "\"")
	timeNow, err := time.Parse(time.ANSIC, body)
	if err != nil {
		return time.Time{}, err
	}
	return timeNow, nil

}
