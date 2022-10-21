package main

import (
	"fmt"
	"time"
)

// function to made a separator or nothing for the output
func write(what string) {
	switch what {
	case "separator":
		fmt.Println("--------------------------------------------------")
	case "nothing":
		fmt.Println("")
	}
}

// main function
func main() {
	state(false)
	selection()
}

// function to select what the user want to do
func selection() {
	var choice string

	fmt.Println("Select an option:")
	fmt.Println("1. Leds Blink")
	fmt.Println("2. Leds On")
	fmt.Println("3. Leds Off")
	fmt.Println("4. Schedule Leds")
	fmt.Println("5. Exit Program")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		blink()
	case "2":
		write("separator")
		state(true)
		write("separator")
		selection()
	case "3":
		write("separator")
		state(false)
		write("separator")
		selection()
	case "4":
		scheduler()
	case "5":
		fmt.Println("Exiting Program")
		return
	}
}

// function to set the sleep time between each led state
func sleep() {
	rate := time.Second / 5
	time.Sleep(rate)
}

// function to know if the leds are on or off
func state(true bool) {
	if true {
		fmt.Println("LED ON")
	} else {
		fmt.Println("LED OFF")
	}
}

func blink() {
	for {
		state(true)
		sleep()
		state(false)
		sleep()
	}
}

func scheduler() {
	var hourOff int
	var minuteOff int
	var hourOn int
	var minuteOn int

	fmt.Println("Enter the hour to turn off the leds:")
	fmt.Scanln(&hourOff)
	write("nothing")
	fmt.Println("Enter the minute to turn off the leds:")
	fmt.Scanln(&minuteOff)
	write("nothing")
	fmt.Println("Enter the hour to turn on the leds:")
	fmt.Scanln(&hourOn)
	write("nothing")
	fmt.Println("Enter the minute to turn on the leds:")
	fmt.Scanln(&minuteOn)
	write("nothing")

	scheduledOperation(hourOff, minuteOff, hourOn, minuteOn)
	fmt.Println("Leds will be turned off at", hourOff, ":", minuteOff, "and turned on at", hourOn, ":", minuteOn)
	write("separator")

}

// Schedule the program to turn Off the leds at a spedific time
func scheduledOperation(hourOff int, minuteOff int, hourOn int, minuteOn int) {

	now := time.Now()
	hr, min, sec := now.Clock()

	if hr == hourOff && min >= minuteOff && sec >= 0 {
		state(false)
	} else {
		fmt.Printf("[%d]hour : [%d]minutes : [%d]seconds\n", hr, min, sec)
		time.Sleep(time.Second * 10)
		scheduledOperation(hourOff, minuteOff, hourOn, minuteOn)
	}

	if hr == hourOn && min >= minuteOn && sec >= 0 {
		state(true)
	} else {
		fmt.Printf("[%d]hour : [%d]minutes : [%d]seconds\n", hr, min, sec)
		time.Sleep(time.Second * 10)
		scheduledOperation(hourOff, minuteOff, hourOn, minuteOn)
	}
}
