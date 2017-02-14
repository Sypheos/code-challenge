package main

import (
	"code-challenge"
	"flag"
	"fmt"
	"os"
)

var uri string

func main() {

	flag.Parse()
	if uri == "" {
		fmt.Println("uri must be specified with option -uri=; see options with --help")
		os.Exit(1)
	}
	code_challenge.Start(uri)
}

func init() {
	flag.StringVar(&uri, "uri", "", "URI to MQTT broker ex: tcp://localhost:1883")
}
