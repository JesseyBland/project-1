package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	// Reverse Proxy points from the load balancer to the reverse proxy
	fmt.Println("Reverse Proxy Initializing...")

	proxy := httputil.NewSingleHostReverseProxy()
	http.ListenAndServe(":8081", proxy)

}
