package cmd

import (
	"flag"
	"net/http"
	"fmt"

	"go.zoe.im/gopkg/gopkg"
)

var (
	// TODO: add https
	addr = flag.String("http", ":8080", "Serve HTTP at given address")
	index = flag.String("index", "https://zoe.im", "Redirect while attach index")
)

func init() {

}

func Run() error {
	flag.Parse()

	opts := []gopkg.Option{gopkg.Index(*index)}
	for _, arg := range flag.Args() {
		opts = append(opts, gopkg.Prefix(arg))
	}

	inst := gopkg.NewInstance(opts...)

	fmt.Printf("[gopkg] going to start server: %s\n", *addr)
	return http.ListenAndServe(*addr, inst)
}
