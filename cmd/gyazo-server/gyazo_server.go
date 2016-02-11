package main

import (
	"github.com/moznion/gyazo-server"
	"os"
)

func main() {
	server.Run(os.Args[1:])
}
