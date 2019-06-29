package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Start() {
	url, err := url.Parse("http://192.168.20.101")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		ip := "127.0.0.1"
		switch req.Host {
		case "192.168.20.109":
			ip = "192.168.20.101"
		case "192.168.20.110":
			ip = "192.168.20.102"
		}
		if req.URL.Path == "/api.php" {
			fmt.Println(ip)
		}
		req.Host = ip
		req.URL.Host = ip
		req.URL.Scheme = "http"
		// Check req.Host
		// Change req.URL
		// Change matrix
	}
	http.ListenAndServe(":9090", proxy)
}
