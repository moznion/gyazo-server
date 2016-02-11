package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/moznion/gyazo-server/aws"
	"github.com/moznion/gyazo-server/controller"
)

const version = "0.0.1"

type opts struct {
	Port       int    `short:"p" long:"port" default:"9090" description:"Port number for listening"`
	BucketName string `short:"b" long:"bucket" required:"true" description:"Bucket name for AWS"`
	Region     string `short:"r" long:"region" required:"true" description:"Region name for AWS"`
}

func parseArgs(args []string) (opt *opts) {
	o := &opts{}
	p := flags.NewParser(o, flags.Default)
	p.Usage = fmt.Sprintf("\n\nVerion:\n  %s", version)
	_, err := p.ParseArgs(args)
	if err != nil {
		os.Exit(1)
	}
	return o
}

func Run(args []string) {
	o := parseArgs(args)

	c := controller.NewController(aws.NewS3Info(o.Region, o.BucketName))

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/{key}":     c.GetImage,            // GET
		"/app/image": c.PostImageFromClient, // POST
	}

	r := mux.NewRouter()
	for p, h := range routes {
		r.HandleFunc(p, h)
	}

	log.Printf("Listen - 127.0.0.1:%d", o.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", o.Port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
