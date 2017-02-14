package code_challenge

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"sort"
)

const TEMP_TOPIC = "/readings/temperature"

var heatHist = make([]float64, 10) //should later be replace by db

const TEMP_LOW = 18
const TEMP_AVG = 20
const TEMP_TARGET = 22
const TEMP_HOT = 25
const TEMP_HELL = 30

var heatMap map[int]float64 = map[int]float64{
	TEMP_LOW:    100,
	TEMP_AVG:    60,
	TEMP_TARGET: 30,
	TEMP_HOT:    5,
	TEMP_HELL:   0}

//startHeat Subscribe to the temperature topic and create a handler.
//The created handler will, if the received message well formatted, call the calcHeatOutput function
func startHeat(client mqtt.Client) {

	for i := range heatHist {
		heatHist[i] = 0
	}
	client.Subscribe(TEMP_TOPIC, QOS, func(client mqtt.Client, msg mqtt.Message) {

		log.Println(msg.Payload())
		m := Sensor{}
		if err := json.Unmarshal(msg.Payload(), &m); err != nil {
			log.Println(err)
			return
		}
		if m.Type == "temperature" {
			if v, ok := m.Value.(float64); ok {
				calcHeatOutput(client, v)
			} else {
				log.Println(TEMP_TOPIC, "Wrong argument type for value in expected float64")
			}
		} else {
			log.Println("wrong sensor type")
		}
	})
}

//stopHeat close the valve and unsuscribe from the topic.
func stopHeat(client mqtt.Client) {

	publishTemp(client, 0)
	client.Unsubscribe(TEMP_TOPIC)
}

//caclHeatOutput calculate the avg temperature based on the last 10 temperature received.
//It will then select the valve output from the heatMap and publish it.
//The algorithm try to make use the heater in phase: Strong, low, strong. in order to save power
func calcHeatOutput(client mqtt.Client, temp float64) {

	heatHist = append(heatHist[1:], temp)

	avg := 0.0
	for _, val := range heatHist {
		avg += float64(val)
	}
	avg = avg / float64(len(heatHist))
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

//publishTemp send val as the valve output on the room1 topic.
//for safety purpose val is maintained between [0 - 100]
func publishTemp(client mqtt.Client, val float64) {
	s, err := json.Marshal(map[string]int{"level": int(val) % 100})
	if err != nil {
		log.Println(err)
		return
	}
	client.Publish(ROOM1_TOPIC, QOS, false, s)
}
