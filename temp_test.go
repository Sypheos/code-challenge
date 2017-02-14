package code_challenge

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var test_uri string = "tcp://localhost:1883"

func TestTempSub(t *testing.T) {

	client := NewClient(test_uri)
	defer client.Disconnect(250)
	startHeat(client)
	ch := make(chan int)
	client.Subscribe(ROOM1_TOPIC, QOS, func(client mqtt.Client, msg mqtt.Message) {

		m := map[string]int{}
		json.Unmarshal(msg.Payload(), &m)
		assert.Equal(t, 10, m["level"])
		ch <- 0
	})
	publishTemp(client, 10)
	select {
	case <-ch:
		return
	case <-time.After(time.Second * 10):
		t.Fail()
	}
}
