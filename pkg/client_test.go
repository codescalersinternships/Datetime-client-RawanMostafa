package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func assertEquality(t *testing.T, obj1 any, obj2 any) {
	t.Helper()
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		t.Errorf("Error! type mismatch, wanted: %t got: %t", reflect.TypeOf(obj1), reflect.TypeOf(obj2))
	}
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Error! values mismatch, wanted: %v got: %v", obj1, obj2)
	}
}

func readBody(t *testing.T, resp *http.Response) ([]byte, error) {
	t.Helper()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), fmt.Errorf("error in reading request body: %v", err)

	}
	return body, nil
}

func TestSendRequest(t *testing.T) {
	formattedTime := time.Now().Format("2024/09/19 12:57:04")
	timeJson, err := json.Marshal(formattedTime)
	if err != nil {
		t.Errorf("error converting to json: %v", err)
	}
	testcases := []struct {
		name        string
		baseUrl     string
		endpoint    string
		port        string
		timeout     time.Duration
		expected    any
		contentType string
		statusCode  int
	}{
		{
			name:        "correct configs, gin, plain text",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "text/plain",
			timeout:     time.Second,
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, gin, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "application/json",
			timeout:     time.Second,
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, plain text",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "text/plain",
			timeout:     time.Second,
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "application/json",
			timeout:     time.Second,
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "unsupported content type",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "text/javascript; charset=utf-8",
			timeout:     time.Second,
			expected:    http.StatusText(http.StatusUnsupportedMediaType),
			statusCode:  http.StatusUnsupportedMediaType,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			c := NewClient(testcase.baseUrl, testcase.endpoint, testcase.port, time.Second)
			resp, err := c.SendRequest(testcase.contentType)
			if err != nil {
				t.Error(err)
			}
			resBody, err := readBody(t, resp)
			if err != nil {
				t.Error(err)
			}

			if testcase.contentType == "application/json" {
				assertEquality(t, testcase.expected, resBody)
			} else {
				assertEquality(t, testcase.expected, string(resBody))
			}
			assertEquality(t, testcase.statusCode, resp.StatusCode)

		})
	}
}
