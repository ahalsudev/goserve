package main

import (
	"flag"
	"fmt"
	"goserve/router"
	"net"
	"net/http"
)

func main() {
	var (
		host = flag.String("host", "", "Host http address to listen on")
		port = flag.String("port", "8000", "Port number for http listener")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	s := http.Server{
		Addr:    addr,
		Handler: router.NewRouter(),
	}
	fmt.Printf("Starting HTTP listener at %s\n", addr)
	s.ListenAndServe()
}
