package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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

func TestRetrySendRequest(t *testing.T) {
	formattedTime := time.Now().Format(time.ANSIC)
	timeJson, err := json.Marshal(formattedTime)
	if err != nil {
		t.Errorf("error converting to json: %v", err)
	}
	testcases := []struct {
		name        string
		baseUrl     string
		endpoint    string
		port        string
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
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, gin, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, plain text",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "text/plain",
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "unsupported content type",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "text/javascript; charset=utf-8",
			expected:    http.StatusText(http.StatusUnsupportedMediaType),
			statusCode:  http.StatusUnsupportedMediaType,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			c := NewClient(testcase.baseUrl, testcase.endpoint, testcase.port, time.Second)
			resp, err := c.RetrySendRequest(testcase.contentType)
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

func TestSendRequest(t *testing.T) {
	formattedTime := time.Now().Format(time.ANSIC)
	timeJson, err := json.Marshal(formattedTime)
	if err != nil {
		t.Errorf("error converting to json: %v", err)
	}
	testcases := []struct {
		name        string
		baseUrl     string
		endpoint    string
		port        string
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
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, gin, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, plain text",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8080",
			contentType: "text/plain",
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, json",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "unsupported content type",
			baseUrl:     "http://localhost",
			endpoint:    "/datetime",
			port:        "8083",
			contentType: "text/javascript; charset=utf-8",
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

func TestWithMockServer(t *testing.T) {
	formattedTime := time.Now().Format(time.ANSIC)
	timeJson, err := json.Marshal(formattedTime)
	if err != nil {
		t.Errorf("error converting to json: %v", err)
	}
	testcases := []struct {
		name        string
		expected    any
		contentType string
		statusCode  int
	}{
		{
			name:        "correct configs, gin, plain text",
			contentType: "text/plain",
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, gin, json",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, plain text",
			contentType: "text/plain",
			expected:    formattedTime,
			statusCode:  http.StatusOK,
		},
		{
			name:        "correct configs, http, json",
			contentType: "application/json",
			expected:    timeJson,
			statusCode:  http.StatusOK,
		},
		{
			name:        "unsupported content type",
			contentType: "text/javascript; charset=utf-8",
			expected:    http.StatusText(http.StatusUnsupportedMediaType) + "\n",
			statusCode:  http.StatusUnsupportedMediaType,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				currentTime := time.Now()
				formattedTime := currentTime.Format(time.ANSIC)

				if strings.Contains(r.Header.Get("content-type"), "text/plain") {

					w.Header().Set("Content-Type", "text/plain")
					fmt.Fprint(w, formattedTime)

				} else if strings.Contains(r.Header.Get("content-type"), "application/json") {

					w.Header().Set("Content-Type", "application/json")

					timeJson, err := json.Marshal(formattedTime)
					if err != nil {
						log.Fatalf("error converting to json: %v", err)
					}
					_, err = w.Write(timeJson)
					if err != nil {
						log.Fatalf("error writing data to response: %v", err)
					}
				} else {
					http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
				}
			}))

			defer mockServer.Close()

			parts := strings.Split(mockServer.URL, ":")
			port := parts[len(parts)-1]

			c := NewClient("http://127.0.0.1", "", port, time.Second)
			resp, err := c.RetrySendRequest(testcase.contentType)
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
