package main

import (
	"fmt"
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
	defer client.Disconnect(250)

	for i := 0; i < num; i++ {
		fmt.Println("---- doing publish ----")
		token := client.Publish(topic, qos, false, "Hello")
		token.Wait()
	}
}
