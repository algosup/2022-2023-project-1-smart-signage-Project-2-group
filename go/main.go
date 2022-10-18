package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	beginning := light_on(true)
	fmt.Println(beginning)
}

func light_on(val bool) bool {
	rate := time.Second / 500

	leds := machine.PA12

	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})

	if val == true {
		for {
			leds.High()
			time.Sleep(rate)
			leds.Low()
			time.Sleep(rate)
		}
		return val
	}
}
