package main

import (
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/rpi/sensor/pico"
)

func main() {
	pico := pico.NewSensor()

	err := pico.Boot()
	if err != nil {
		log.Fatalf("could not boot pico sensor software: %v", err)
	}

	fmt.Println("Hello Pico!")
}
