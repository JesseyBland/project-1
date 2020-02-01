package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// NewMultipleHostReverseProxy creates a reverse proxy
func NewMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {

	director := func(req *http.Request) {
		println("CALLING DIRECTOR")

		target := loadBal()

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}

	}
	return &httputil.ReverseProxy{
		Director: director,
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				println("CALLING PROXY")
				return http.ProxyFromEnvironment(req)
			},
			Dial: func(network, addr string) (net.Conn, error) {
				println("CALLING DIAL")
				conn, err := (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial(network, addr)
				if err != nil {
					println("Error during DIAL:", err.Error())
				}
				return conn, err
			},
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func main() {
	targets := []*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:8888",
		},
		{
			Scheme: "http",
			Host:   "localhost:7777",
		},
		{
			Scheme: "http",
			Host:   "localhost:9999",
		},
	}

	fmt.Println(" RProxyHost:Localhost | Port: 9090")

	proxy := NewMultipleHostReverseProxy(targets)

	http.ListenAndServe(":9090", proxy)

}

var a, b, c int

func loadBal() *url.URL {

	if a == b && a == c {

		a++

		fmt.Println("On Server :8888   a=", a)
		ticURL, err := url.Parse("http://localhost:8888")

		if err != nil {
			log.Fatal(err)
		}
		return ticURL

	} else if a > b && b == c {
		b++
		fmt.Println("On Server :9999   b=", b)
		ticURL, err1 := url.Parse("http://localhost:9999")
		if err1 != nil {
			log.Fatal(err1)
		}
		return ticURL

	} else if a > c && b > c {
		c++
		fmt.Println("On Server :7777   c=", c)
		ticURL, err2 := url.Parse("http://localhost:7777")
		if err2 != nil {
			log.Fatal(err2)
		}
		return ticURL
	}

	return nil

}
