package main

import (
	"machine"
	// "fmt"
	"time"
)

func main() {
	sendMsg("petit test")
}

func Off() {
	leds := machine.PA12
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.Low()
}

func On() {
	leds := machine.PA12
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.High()
}

func minLight() {
	for {
		On()
		time.Sleep(time.Second / 1000)
		Off()
		time.Sleep(time.Second / 100)
	}
}

func sendMsg(message string) {
	message = "\"" + message + "\"" + "\r\n"
	request := "AT+MSG=" + message

	uart := machine.UART2
	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})
	_, err := uart.Write([]byte("AT+JOIN\r\n"))

	if err != nil {
		println("Error: " + err.Error())

		// return ""
	}
	msg1 := ""
	for {
		if uart.Buffered() > 0 {
			rb, err := uart.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				continue
				// return ""
			}
			msg1 += string(rb)
			if msg1[len(msg1)-1] == '\n' {
				break
			}
		}
		uart.Write([]byte(request))
	}

	// fmt.Println(request)
}
