package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/emilieanthony/senzr/rpi/cloud"
	"github.com/emilieanthony/senzr/rpi/sensor/pico"
)

const (
	readInterval = 5 * time.Second
)

// TODO: cloud logger?
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// init pubsub client
	pubSubClient, err := cloud.NewPubSubClient(ctx)
	if err != nil {
		log.Fatalf("pubsub: %v", err)
	}
	defer func() {
		pubSubClient.Stop()
		cancel()
	}()
	// init sensor
	sensor := pico.NewSensor()
	fmt.Printf("Successfully booted application! \n")
	sensorDataChannel := make(chan *pico.Data)
	// start sensor read polling
	go poll(readInterval, func() {
		var data pico.Data
		if err := sensor.Read(ctx, &data); err != nil {
			fmt.Printf("error reading sensor data: %v", err)
		} else {
			sensorDataChannel <- &data
		}
	})

	for data := range sensorDataChannel {
		pubSubClient.Publish(ctx, data)
	}
}

func poll(interval time.Duration, fn func()) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		fn()
	}
}
