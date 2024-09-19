package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func CreateClient() http.Client {
	return http.Client{Timeout: time.Duration(1) * time.Second}
}

func ReadBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading request body: %v", err)

	}
	return string(body), nil
}

func SendRequest(c http.Client) (*http.Response, error) {
	url := os.Getenv("DATETIME_URL")
	endpoint := os.Getenv("DATETIME_ENDPOINT")
	port := os.Getenv("DATETIME_PORT")

	if url == "" {
		url = "http://localhost"
	}
	if endpoint == "" {
		endpoint = "/datetime"
	}
	if port == "" {
		port = "8083"
	}

	resp, err := c.Get(url + ":" + port + endpoint)
	if err != nil {
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	return resp, nil
}
