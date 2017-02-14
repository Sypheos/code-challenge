package code_challenge

import (
	"bufio"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

const ROOM1_TOPIC = "/actuators/room-1"
const QOS byte = 0

type Sensor struct {
	SensorID string      `json:"sensor_id"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
}

func Start(uri string) {

	client := NewClient(uri)
	defer client.Disconnect(250)

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 1000):
				reader := bufio.NewReader(os.Stdin)
				text, _ := reader.ReadString('\n')
				if text == "quit\n" {
					return
				} else {
					fmt.Println("Type \"quit\" to exit")
				}
			}
		}
	}()

	ch := startMotion(client)
	for {
		select {
		case b := <-ch:
			log.Println("state changed")
			if b {
				startHeat(client)
			} else {
				stopHeat(client)
			}

		}
	}
}

func NewClient(uri string) (client mqtt.Client) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
