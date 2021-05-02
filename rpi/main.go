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
		log.Fatal("could not boot pico sensor software: %w", err)
	}

	fmt.Println("Hello Pico!")
}
