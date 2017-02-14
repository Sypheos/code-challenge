package main

import (
	"fmt"
	"time"
	"flag"
	"bufio"
	"os"
)

var uri string

func main() {

	flag.Parse()
	if uri == "" {
		fmt.Println("uri must be specified with option -uri=; see options with --help")
		os.Exit(1)
	}
	client := newClient(uri)
	defer client.Disconnect(250)

	registerTemp(client)
	for {
		select {
		case <-time.After(time.Millisecond * 500):
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			if text == "quit\n" {
				return
			} else {
				fmt.Println("Type \"quit\" to exit")
			}
		}
	}
}

func init() {
	flag.StringVar(&uri, "uri", "", "URI to MQTT broker ex: tcp://localhost:1883")
}