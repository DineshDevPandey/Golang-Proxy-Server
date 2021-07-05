package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Args struct {
	key       string
	value     string
	serverURL string
}

func (arg *Args) Handler(rw http.ResponseWriter, req *http.Request) {

	// prepare request object for actual API
	serverURL, err := url.Parse(arg.serverURL)
	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		fmt.Fprint(rw, err)
		return
	}

	req.Host = serverURL.Host
	req.URL.Host = serverURL.Host
	req.URL.Scheme = serverURL.Scheme

	// Stop all other paths except "/test"
	// We know endpoint of one API only
	// If we know server URL address, no need to do this
	if req.URL.Path != "/test" {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, "Path not found, Try http://127.0.0.1:PORT_NUMBER/test")
		return
	}

	req.URL.Path = serverURL.Path

	// http: Request.RequestURI can't be set in client requests
	req.RequestURI = ""

	// set X-Forwarded-For so that server can receive clients IP
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	req.Header.Set("X-Forwarded-For", host)
	req.Header.Set("X-Forwarded-Proto", "http")

	// setting my secret key and value
	if len(arg.key) > 0 {
		req.Header.Set("Key", "the-secret-key")
	}

	// make request to server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}

	// copy header from server response to proxy response
	for key, values := range resp.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}

	// for streaming data we need to flush iterativally till the complete data is received from server
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.Tick(10 * time.Millisecond):
				rw.(http.Flusher).Flush()
			case <-done:
				return
			}
		}
	}()

	// if server sends trailer, we need to pass it to client

	// announce trailer keys
	trailerKeys := []string{}
	for key := range resp.Trailer {
		trailerKeys = append(trailerKeys, key)
	}
	rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

	// copy header StatusCode and body from server response to proxy response
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)

	// after reading body fill the trailer value
	for key, values := range resp.Trailer {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}

	close(done)
}

func Port() string {
	// default port is 8080
	// if received port is not a numeric value server will start on default port
	ports := os.Args
	port := ":8080"

	if len(os.Args) >= 2 && len(ports[1]) > 0 {
		if _, err := strconv.Atoi(ports[1]); err == nil {
			port = ":" + ports[1]
		}
	}
	return port
}

func main() {

	arg := Args{"Key", "the-secret-key", "https://postman-echo.com/headers"}
	proxy := http.HandlerFunc(arg.Handler)

	port := Port()
	log.Println("Starting proxy server on : http://127.0.0.1", port)

	if err := http.ListenAndServe(port, proxy); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
