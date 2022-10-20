package main

import "testing"

func TestBlink(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		want := true
		got := l.Set(want)
		if got != want {
			t.Errorf("The light isn't turn on")
		}
	})
}
