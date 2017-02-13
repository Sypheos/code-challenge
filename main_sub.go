package main

import (
	"fmt"
	"os"
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var topic string = "test"
var qos byte = 0
var num int = 5

func main() {

	choke := make(chan [2]string)

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetDefaultPublishHandler(func(client *mqtt.Client, msg mqtt.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	if token := client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for receiveCount := 0; receiveCount < num; receiveCount++ {
		incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	}
	client.Disconnect(250)
}