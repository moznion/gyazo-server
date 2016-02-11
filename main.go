package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/jessevdk/go-flags"
	"github.com/moznion/gyazo-server/service"
)

const version = "0.0.1"

type opts struct {
	Port int `short:"p" long:"port" default:"9090" description:"Port number for listening"`
}

func parseArgs(args []string) (opt *opts) {
	o := &opts{}
	p := flags.NewParser(o, flags.Default)
	p.Usage = fmt.Sprintf("\n\nVerion:\n  %s", version)
	p.ParseArgs(args)
	return o
}

func Run(args []string) {
	o := parseArgs(args)

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/app/image": service.PostImageFromClient, // POST
	}

	s := http.NewServeMux()
	for p, h := range routes {
		s.Handle(p, handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h)))
	}

	log.Printf("Listen - 127.0.0.1:%d", o.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", o.Port), handlers.CompressHandler(s))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
