package pkg

import (
	"fmt"
	"io"
	"net/http"
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

func SendRequest(c http.Client, port string) (*http.Response, error) {
	resp, err := c.Get("http://localhost:" + port + "/datetime")
	if err != nil {
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	return resp, nil
}
