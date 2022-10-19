package main

import "testing"

func TestBlinkLight(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		want := true
		got := blinkLight(want)
		if got != want {
			t.Errorf("The light isn't turn on")
		}
	})
}
