package pico

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	processName = "/usr/local/bin/usbtenkiget"
	/*
		the -i flag takes comma-separated sensor channels.
		00 - co2
		01 - temperature
		02 - humidity
		source: https://www.dracal.com/store/products/usb_dxc100/index.php
	*/
	channelsFlag            = "-i 00,01,02"
	co2ChannelIndex         = 0
	temperatureChannelIndex = 1
	humidityChannelIndex    = 2
)

type Data struct {
	Co2         float64   `json:"co2"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
}

func (d *Data) Encode() ([]byte, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type Sensor struct{}

func NewSensor() *Sensor {
	return &Sensor{}
}

/*
	starts child process that reads sensor data and keeps reader and writer references to that process
	source: https://www.dracal.com/store/support/programmers_howto/index.php
*/
func (s *Sensor) ReadValue(ctx context.Context) (*bytes.Buffer, error) {
	var data bytes.Buffer
	process := exec.CommandContext(ctx, processName, channelsFlag)
	process.Stdout = &data
	err := process.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start process %s: %w", processName, err)
	}
	return &data, nil
}

func (s *Sensor) Read(ctx context.Context, d *Data) error {
	now := time.Now().Format(time.RFC3339)
	data, err := s.ReadValue(ctx)
	if err != nil {
		return fmt.Errorf("reading value: %w", err)
	}
	fmt.Printf("[%s] read %d bytes \n", now, data.Len())
	rawData := strings.Split(data.String(), ",")
	if len(rawData) > co2ChannelIndex {
		f, err := strconv.ParseFloat(rawData[co2ChannelIndex], 32)
		// ignore parse error and just skip saving sensor data
		if err == nil {
			d.Co2 = f
		}
	}
	if len(rawData) > temperatureChannelIndex {
		f, err := strconv.ParseFloat(strings.TrimSpace(rawData[temperatureChannelIndex]), 32)
		// ignore parse error and just skip saving sensor data
		if err != nil {
			fmt.Printf("warning: temperature parse error: %v \n", err)
		}
		if err == nil {
			d.Temperature = f
		}
	}
	if len(rawData) > humidityChannelIndex {
		f, err := strconv.ParseFloat(strings.TrimSpace(rawData[humidityChannelIndex]), 32)
		// ignore parse error and just skip saving sensor data
		if err != nil {
			fmt.Printf("warning: humidity parse error: %v \n", err)
		}
		if err == nil {
			d.Humidity = f
		}
	}
	d.Timestamp = time.Now()
	return nil
}
