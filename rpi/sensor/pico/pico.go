package pico

import (
	"fmt"
	"io"
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
	Co2         float64
	Temperature float64
	Humidity    float64
}

type Sensor struct {
	process *exec.Cmd
}

func NewSensor() *Sensor {
	return &Sensor{}
}

/*
	starts child process that reads sensor data and keeps reader and writer references to that process
	source: https://www.dracal.com/store/support/programmers_howto/index.php
*/
func (s *Sensor) Boot() error {
	process := exec.Command(processName, channelsFlag)
	err := process.Start()
	if err != nil {
		return fmt.Errorf("could not start process %s: %w", processName, err)
	}
	err = process.Wait() // wait for process to start
	if err != nil {
		return fmt.Errorf("failed waiting for process %s: %w", processName, err)
	}

	if process.ProcessState.Exited() {
		return fmt.Errorf("process %s failed unexpectedly", processName)
	}
	s.process = process
	return nil
}

func (s *Sensor) ProcessExited() bool {
	return s.process.ProcessState.Exited()
}

func (s *Sensor) Read(d *Data) error {
	if s.ProcessExited() {
		return fmt.Errorf("process %s is not running", processName)
	}
	reader := io.Reader(s.process.Stdin)
	b := make([]byte, 0)
	now := time.Now().String()
	n, err := reader.Read(b)
	if err != nil {
		return fmt.Errorf("[%s] could not read data: %w", now, err)
	}
	fmt.Printf("[%s] read %d bytes", now, n)
	rawData := strings.Split(string(b), ",")
	if len(rawData) > co2ChannelIndex+1 {
		f, err := strconv.ParseFloat(rawData[co2ChannelIndex], 32)
		// ignore parse error and just skip saving sensor data
		if err == nil {
			d.Co2 = f
		}
	}
	if len(rawData) > temperatureChannelIndex+1 {
		f, err := strconv.ParseFloat(rawData[temperatureChannelIndex], 32)
		// ignore parse error and just skip saving sensor data
		if err == nil {
			d.Temperature = f
		}
	}
	if len(rawData) > humidityChannelIndex+1 {
		f, err := strconv.ParseFloat(rawData[humidityChannelIndex], 32)
		// ignore parse error and just skip saving sensor data
		if err == nil {
			d.Humidity = f
		}
	}
	return nil
}
