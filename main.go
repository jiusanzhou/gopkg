package main

import (
	"log"

	"go.zoe.im/gopkg/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
