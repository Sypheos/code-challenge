package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"encoding/json"
)

const MOTION_TOPIC = "/readings/motion"
var state bool = false

//startMotion create a subscription to the motion topic.
//return a channel where change of activity are pushed.
//true for activity, false for  none.
func startMotion(client mqtt.Client) (ch chan bool) {

	for i := range heatHist {
		heatHist[i] = 0
	}
	client.Subscribe(MOTION_TOPIC, QOS, func(client mqtt.Client, msg mqtt.Message) {

		m := sensor{}
		if err := json.Unmarshal(msg.Payload(), &m); err != nil {
			log.Println(err)
			return
		}
		if m.Type == "motion" {
			if b, ok := m.Value.(bool); ok && b != state {
				ch <- b
			} else {
				log.Println(MOTION_TOPIC, "Wrong argument type for value in expected boolean")
			}
		}
	})
	return ch
}