package main

import (
	"fmt"
	"time"
)

// main function
func main() {
	scheduler()
}

// function to know if the leds are on or off
func state(bool) {
	if true {
		fmt.Println("LED ON")
	} else if false {
		fmt.Println("LED OFF")
	}
}

func scheduler() {
	var hourOff int
	var minuteOff int
	var hourOn int
	var minuteOn int

	fmt.Println("Enter the hour to turn off the leds:")
	fmt.Scanln(&hourOff)
	fmt.Println("Enter the minute to turn off the leds:")
	fmt.Scanln(&minuteOff)
	fmt.Println("Enter the hour to turn on the leds:")
	fmt.Scanln(&hourOn)
	fmt.Println("Enter the minute to turn on the leds:")
	fmt.Scanln(&minuteOn)

	scheduledOperation(hourOff, minuteOff, hourOn, minuteOn)
	fmt.Println("Leds will be turned off at", hourOff, ":", minuteOff, "and turned on at", hourOn, ":", minuteOn)
}

func scheduledOperation(hourOff int, minuteOff int, hourOn int, minuteOn int) {
	for {
		now := time.Now()
		hr, min, sec := now.Clock()

		if hr == hourOff && min >= minuteOff && min != minuteOn {
			time.Sleep(time.Second)
			fmt.Println("Time now: ", hr, ":", min, ":", sec)
			state(false)
			scheduledOperation(hourOn, minuteOn, hourOff, minuteOff)
		} else if hr == hourOn && min >= minuteOn && min != minuteOff {
			time.Sleep(time.Second)
			state(true)
			scheduledOperation(hourOff, minuteOff, hourOn, minuteOn)
			fmt.Println("Time now: ", hr, ":", min, ":", sec)
		} else {
			//print to log the time now
			fmt.Println("Time now: ", hr, ":", min, ":", sec)
			time.Sleep(time.Second)
		}
	}
}