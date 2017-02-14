package main

import (
	"log"
	"github.com/eclipse/paho.mqtt.golang"
	"encoding/json"
	"sort"
)

const TEMP_TOPIC = "/readings/temperature"
const ROOM1_TOPIC = "/actuators/room-1"
const QOS byte = 0

var heatHist = make([]float64, 10) //should later be replace by db

const TEMP_LOW = 18
const TEMP_AVG = 20
const TEMP_TARGET = 22
const TEMP_HOT = 25
const TEMP_HELL = 30

var heatMap map[int] float64 = map[int] float64 {
	TEMP_LOW: 100,
	TEMP_AVG: 60,
	TEMP_TARGET: 30,
	TEMP_HOT: 5,
	TEMP_HELL: 0}


type sensor struct {

	SensorID string `json:"sensor_id"`
	Type     string `json:"type"`
	Value    float64 `json:"value"`
}

type home struct {
	mqtt.Client
}

func registerTemp(client mqtt.Client) {

	for i := range heatHist {
		heatHist[i] = 0
	}
	client.Subscribe(TEMP_TOPIC, QOS, func(client mqtt.Client, msg mqtt.Message) {
		log.Println(string(msg.Payload()))

		m := sensor{}
		if err := json.Unmarshal(msg.Payload(), &m); err != nil {
			log.Println(err)
			return
		}
		if m.Type == "temperature" {
			calcHeatOutput(client, m.Value)
		}
	})
}

func calcHeatOutput(client mqtt.Client, temp float64) {

	heatHist = append(heatHist[1:], temp)

	avg := 0.0
	for _, val := range heatHist {
		avg += float64(val)
	}
	avg = avg/float64(len(heatHist))
	keys := []int{}
	for key := range heatMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, val := range keys {
		if avg < heatMap[val] {
			publishTemp(client, heatMap[val])
			break
		}
	}
}

func publishTemp(client mqtt.Client, val float64) {
	s, err := json.Marshal(map[string]float64{"level":val})
	if err != nil {
		log.Println(err)
		return
	}
	client.Publish(ROOM1_TOPIC, QOS, false, s)
}