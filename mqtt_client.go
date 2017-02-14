package main

import (
	"github.com/eclipse/paho.mqtt.golang"
)

func newClient(uri string) (client mqtt.Client) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
