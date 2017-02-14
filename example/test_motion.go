package main

import (
	"code-challenge"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var t code_challenge.Sensor = code_challenge.Sensor{SensorID: "sensor-1", Type: "temperature", Value: float64(15)}
var activity code_challenge.Sensor = code_challenge.Sensor{SensorID: "m1", Type: "motion", Value: bool(true)}

func main() {

	client := code_challenge.NewClient("tcp://localhost:1883")
	client.Subscribe(code_challenge.ROOM1_TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println(string(msg.Payload()))
	})

	str, err := json.Marshal(activity)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		temp, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
			return
		}
		for {
			select {
			case <-time.After(time.Millisecond * 1000):
				client.Publish(code_challenge.TEMP_TOPIC, code_challenge.QOS, false, temp)
			}
		}
	}()
	for {
		select {
		case <-time.After(time.Millisecond * 5000):
			str, err = json.Marshal(activity)
			if err != nil {
				log.Println(err)
				return
			}
			client.Publish(code_challenge.MOTION_TOPIC, 0, false, str)
			if activity.Value == true {
				activity.Value = false
			} else {
				activity.Value = true
			}
			log.Println(activity)
		}
	}
}
