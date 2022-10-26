package realled

import (
	"machine"
	"time"
)

type LED struct {
	pin machine.Pin
	on  bool
}

func New() *LED {
	leds := machine.PA12
	leds.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	l := LED{
		pin: leds,
		on:  false,
	}
}

// On indicates if the LED is turned on.
func (l *LED) On() bool {
	return l.on
}

// Set sets the LED to the value.
// true will turn the LED on
// false will turn the LED off
func (l *LED) Set(value bool) {
	l.on = value
	if l.on {
		l.pin.High()
		return
	} else {
		l.pin.Low()
	}
}

func (l *LED) Blink() {
	rate := time.Second / 500
	for i := 0; i < 1000; i++ {
		// LED.High()
		l.Set(true)
		time.Sleep(rate)
		// LED.Low()
		l.Set(false)
		time.Sleep(rate)
	}
}
