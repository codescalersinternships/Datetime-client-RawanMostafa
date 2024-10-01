package pkg

import (
	"encoding/json"
	"fmt"
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

func mockServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
}

func TestGetTime(t *testing.T) {
	formattedTime := time.Now().Format(time.ANSIC)
	expectedTime, err := time.Parse(time.ANSIC, formattedTime)
	if err != nil {
		t.Errorf("error parsing the time: %v", err)
	}
	testcases := []struct {
		name          string
		baseUrl       string
		endpoint      string
		port          string
		expectedTime  time.Time
		expectedError error
		contentType   string
	}{
		{
			name:          "correct configs, gin, plain text",
			baseUrl:       "http://localhost",
			endpoint:      "/datetime",
			port:          "8083",
			contentType:   "text/plain",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, gin, json",
			baseUrl:       "http://localhost",
			endpoint:      "/datetime",
			port:          "8080",
			contentType:   "application/json",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, http, plain text",
			baseUrl:       "http://localhost",
			endpoint:      "/datetime",
			port:          "8080",
			contentType:   "text/plain",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, http, json",
			baseUrl:       "http://localhost",
			endpoint:      "/datetime",
			port:          "8083",
			contentType:   "application/json",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "unsupported content type",
			baseUrl:       "http://localhost",
			endpoint:      "/datetime",
			port:          "8083",
			contentType:   "text/javascript; charset=utf-8",
			expectedTime:  time.Time{},
			expectedError: fmt.Errorf("%s", http.StatusText(http.StatusUnsupportedMediaType)),
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			c := NewClient(testcase.baseUrl, testcase.endpoint, testcase.port, testcase.contentType, time.Second)
			timeNow, err := c.GetTime()

			assertEquality(t, testcase.expectedTime, timeNow)
			assertEquality(t, testcase.expectedError, err)

		})
	}
}

func TestGetTimeMock(t *testing.T) {
	formattedTime := time.Now().Format(time.ANSIC)
	expectedTime, err := time.Parse(time.ANSIC, formattedTime)
	if err != nil {
		t.Errorf("error parsing the time: %v", err)
	}
	testcases := []struct {
		name          string
		expectedTime  time.Time
		expectedError error
		contentType   string
	}{
		{
			name:          "correct configs, gin, plain text",
			contentType:   "text/plain",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, gin, json",
			contentType:   "application/json",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, http, plain text",
			contentType:   "text/plain",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "correct configs, http, json",
			contentType:   "application/json",
			expectedTime:  expectedTime,
			expectedError: nil,
		},
		{
			name:          "unsupported content type",
			contentType:   "text/javascript; charset=utf-8",
			expectedTime:  time.Time{},
			expectedError: fmt.Errorf("%s", http.StatusText(http.StatusUnsupportedMediaType)),
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockServer := mockServer(t)

			defer mockServer.Close()

			parts := strings.Split(mockServer.URL, ":")
			port := parts[len(parts)-1]

			c := NewClient("http://127.0.0.1", "", port, testcase.contentType, time.Second)
			timeNow, err := c.GetTime()

			assertEquality(t, testcase.expectedTime, timeNow)

			assertEquality(t, testcase.expectedError, err)

		})
	}
}
