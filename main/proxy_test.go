package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const (
	TrueKey   = "key"
	TrueValue = "the-secret-key"
)

func TestHandler(t *testing.T) {
	arg := Args{"key", "the-secret-key", "https://postman-echo.com/headers"}

	tt := []struct {
		name       string
		method     string
		url        string
		remoteAddr string
		key        string
		value      string
		statusCode int
		serverURL  string
	}{
		{"test URL check", "GET", "http://127.0.0.1:3000/test", "127.0.0.1:56192", arg.key, arg.value, http.StatusOK, arg.serverURL},
		{"empty server URL check", "GET", "http://127.0.0.1:3000/test", "127.0.0.1:56192", arg.key, arg.value, http.StatusOK, ""},
		{"without port remoteAddr check", "GET", "http://127.0.0.1:3000/test", "127.0.0.1", "", "", http.StatusInternalServerError, arg.serverURL},
		{"other URL check", "GET", "http://127.0.0.1:3000/test1", "127.0.0.1:56192", "", "", http.StatusNotFound, arg.serverURL},
		{"response header secret KEY/VALUE check", "GET", "http://127.0.0.1:3000/test", "127.0.0.1:56192", arg.key, arg.value, http.StatusOK, arg.serverURL},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Handler function takes two arguments Request and ResponseWriter
			// create request
			req, err := http.NewRequest(tc.method, tc.url, nil)
			if err != nil {
				t.Fatalf("Could not create request : %v", err)
			}

			// for testing initialized RemoteAddr with localhost and some random port
			req.RemoteAddr = tc.remoteAddr

			// create ResponseWriter
			rw := httptest.NewRecorder()

			// // for test case response header secret KEY/VALUE check
			// if tc.key != TrueKey {
			// 	arg.key = tc.key
			// }
			// if tc.value != TrueValue {
			// 	arg.value = tc.value
			// }
			arg.Handler(rw, req)
			res := rw.Result()

			defer res.Body.Close()
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Could not read body : %v", err)
			}

			body := string(bodyBytes)

			// TEST CASE 1: empty response body
			if len(body) == 0 {
				t.Fatalf("response body can not be empty")
			}

			if len(tc.key) > 0 {
				k := fmt.Sprintf("%q", TrueKey)
				v := fmt.Sprintf("%q", TrueValue)

				//  TEST : key value mismatch for secret key
				if !strings.Contains(body, k+":"+v) {
					t.Errorf("expected value for header key : the-secret-key; got different")
				}
			}
			//  TEST CASE 4: status code
			if res.StatusCode != tc.statusCode {
				t.Errorf("expected status 200; got %v", res.StatusCode)
			}
		})
	}
}

func TestPort(t *testing.T) {

	tt := []struct {
		name  string
		port  string
		cport string
	}{
		{"user defined port", "3000", ":3000"},
		{"non integer port", "abcd", ":8080"},
		{"default port", "", ":8080"},
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = []string{"go run proxy.go", tc.port}
			port := Port()
			if port != tc.cport {
				t.Fatalf("expected port %v; got %v", tc.cport, port)
			}
		})
	}
}
