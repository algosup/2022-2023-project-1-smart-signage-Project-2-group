package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	switchLed := machine.PA1
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		switchLed.High()
		time.Sleep(time.Second)
		led.Low()
		switchLed.Low()
		time.Sleep(time.Second)
	}
}
