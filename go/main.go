package main

import (
	"machine"
	"time"
)

func light_on(val string) string {
	rate := time.Second / 500

	leds := machine.PA12

	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})

	if val == "true" {
		for {
			leds.High()
			time.Sleep(rate)
			leds.Low()
			time.Sleep(rate)
		}
		return val
	}
}
