package main

import (
	"machine"
	"time"
)

func main() {
	rate := time.Second / 500

	leds := machine.PA12

	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		leds.High()
		time.Sleep(rate)
		leds.Low()
		time.Sleep(rate)
	}
}
