package main

import (
	"machine"
	"time"
)

func main() {
	for {
		morning()
		earlyAfternoon()
		earlyEvening()
		evening()
		night()
	}
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
	Off()
	time.Sleep(time.Second / 1000)
	On()
	time.Sleep(time.Second / 100)
}

func temperedLight() {
	Off()
	time.Sleep(time.Second / 700)
	On()
	time.Sleep(time.Second / 700)
}

// 25s -> night(1am-6am) -> off
// 20s -> morning(6am-10am) -> 70%
// 35s -> early afternoon(10am-5pm) -> on
// 20s -> early evening(5pm-9pm) -> 70%
// 20s -> evening(9pm-1am) -> 30%

func night() {
	for start := time.Now(); time.Since(start) < time.Second*25; {
		Off()
	}
	sendMsg("night: lights off")
}

func morning() {
	for start := time.Now(); time.Since(start) < time.Second*20; {
		minLight()
	}
	sendMsg("morning: lights on")
}

func earlyAfternoon() {
	for start := time.Now(); time.Since(start) < time.Second*35; {
		On()
	}
	sendMsg("early afternoon: lights on")
}

func earlyEvening() {
	for start := time.Now(); time.Since(start) < time.Second*20; {
		minLight()
	}
	sendMsg("early evening: lights on")
}

func evening() {
	for start := time.Now(); time.Since(start) < time.Second*20; {
		temperedLight()
	}
	sendMsg("evening: lights on")
}

func sendMsg(message string) {
	message = "\"" + message + "\"" + "\r\n"
	request := "AT+MSG=" + message

	uart := machine.UART2
	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})
	_, err := uart.Write([]byte("AT+JOIN\r\n"))

	if err != nil {
		println("Error: " + err.Error())
	}
	msg1 := ""
	for {
		if uart.Buffered() > 0 {
			rb, err := uart.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				continue
			}
			msg1 += string(rb)
			if msg1[len(msg1)-1] == '\n' {
				break
			}
		}
		uart.Write([]byte(request))
	}
}
