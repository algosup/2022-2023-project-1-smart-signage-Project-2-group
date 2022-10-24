package realled

import "testing"

func TestOn(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		l := LED{
			on: true,
		}
		want := false
		got := l.On()
		if got != want {
			t.Errorf("The light isn't turn on")
		}
	})
}

// func TestSet(t *testing.T) {
// 	t.Run("false", func(t *testing.T) {
// 		l := LED{
// 			on: true
// 		}
// 		want := true
// 		got := l.Set(false)
// 		if got != want {
// 			t.Errorf("The light isn't turn on")
// 		}
// 	})
// }

// func TestBlink(t *testing.T) {
// 	t.Run("true", func(t *testing.T) {
// 		want := true
// 		got := l.Blink()
// 		if got != want {
// 			t.Errorf("The light isn't turn on")
// 		}
// 	})
// }HELLO
