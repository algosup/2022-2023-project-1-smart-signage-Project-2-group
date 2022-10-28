package main

import "time"

type LED struct {
	on bool
}

func New() *LED {
	l := LED{
		on: false,
	}
	return &l
}

// On indicates if the LED is turned on.
func (l *LED) On() bool {
	return l.on
}

// Set sets the LED to the value.
// true will turn the LED on
// false will turn the LED off
func (l *LED) Set(value bool) bool {
	l.on = value
	if l.on {
		return true
	} else {
		return false
	}
}

func (l *LED) Blink() {
	rate := time.Second / 500
	for i := 0; i < 10; i++ {
		l.Set(true)
		time.Sleep(rate)
		l.Set(false)
		time.Sleep(rate)
	}
}
