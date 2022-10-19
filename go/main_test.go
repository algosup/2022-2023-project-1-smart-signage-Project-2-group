package main

import "testing"

func TestLight_on(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		want := true
		got := light_on(want)
		if got != want {
			t.Errorf("The light isn't turn on")
		}
	})
}
