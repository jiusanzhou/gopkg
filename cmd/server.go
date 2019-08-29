package cmd

import (
	"fmt"
	"net/http"
)

type server struct {
	addr  string
	index string
}

func newServer() *server {
	return &server{}
}

// ServeHTTP implement http.Handler
func (svr *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func (svr *server) run() error {
	fmt.Printf("[gopkg] going to start server: %s\n", svr.addr)
	return http.ListenAndServe(svr.addr, svr)
}
