package main

import (
	"net"
	"net/http"

	"github.com/HAL-CaseStudy-1/dev-server/controllers"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/top", controllers.TopHandler)

	if err := http.Serve(listener, nil); err != nil {
		panic(err)
	}
}
