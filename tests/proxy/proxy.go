package proxy

import (
	"fmt"
	"net"
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
		port := 15
		switch req.Host {
		case "192.168.20.201":
			ip = "192.168.20.101"
			port = 15
		case "192.168.20.202":
			ip = "192.168.20.103"
			port = 5
		case "192.168.20.203":
			ip = "192.168.20.102"
			port = 16
		case "192.168.20.204":
			ip = "192.168.20.104"
			port = 1
		case "192.168.20.205":
			ip = "192.168.20.105"
			port = 7
		case "192.168.20.206":
			ip = "192.168.20.106"
			port = 2
		case "192.168.20.207":
			ip = "192.168.20.107"
			port = 3
		}
		if req.URL.Path == "/api.php" {
			fmt.Println(ip)
			go swithMatrix(port)
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

func swithMatrix(port int) {
	outputPort := 1
	conn, err := net.Dial("tcp", "192.168.20.242:9990")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	payload := fmt.Sprintf("VIDEO OUTPUT ROUTING:\n%v %v\n\n", outputPort-1, port-1)
	fmt.Println(payload)

	conn.Write([]byte(payload))

}
