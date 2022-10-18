package main

import "testing"

func TestLight_on(t *testing.T) {
	t.Run(true, func(t *testing.T) {
		val := true
		got := light_on(val)
		if got != true {
			t.Error("The light isn't turn on")
		}
	})
}
