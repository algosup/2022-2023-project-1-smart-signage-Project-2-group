package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	val := true
	beginning := blinkLight(val)
	fmt.Println(beginning)
}

func blinkLight(val bool) bool {
	rate := time.Second / 500

	leds := machine.PA12

	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})

	if val == true {
		for i := 0; i < 1000; i++ {
			leds.High()
			time.Sleep(rate)
			leds.Low()
			time.Sleep(rate)
		}
		return val
	}
	return !val
}
