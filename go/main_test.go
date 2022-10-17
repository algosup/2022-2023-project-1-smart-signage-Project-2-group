package main

import (
	"fmt"
	"testing"
)

func TestLight_on(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		val := "true"
		got := light_on(val)
		if got != "true" {
			fmt.Println("The light isn't turn on")
		}
	})
}

// func TestLight_on(t *testing.T) {
// 	t.Run("false", func(t *testing.T) {
// 		val := "false"
// 		got := TestLight_on(val)
// 		if got == "false" {
// 			fmt.Println("The light is turn off")
// 		}
// 	})
// }
