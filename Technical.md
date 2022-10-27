# Technical Specification
 2022-2023 Project Smart Signage Groupe 2
 
  Team:
  
- [Rémy Charles](https://github.com/RemyCHARLES)
- [Pierre Gorin](https://github.com/Pierre2103)
- [Élise Gautier](https://github.com/elisegtr)
- [Salaheddine Namir](https://github.com/T3rryc)
- [Grégory Pagnoux](https://github.com/Gregory-Pagnoux)

Create on Sep 29, 2022

Last Update Oct 25, 2022
## Table of contents


- [1. Introduction](#1-introduction)
  - [A. Overview](#a-overview)
  - [B. Glossary](#b-glossary)
- [2. Development tools](#2-development-tools)
  - [A. Software used](#a-software-used)
    - [a. Technologies used](#a-technologies-used)
    - [b. Language used](#b-language-used)
  - [B. Hardware used](#b-hardware-used)
- [3. Developement Process](#3-developement-process)
  - [A. microcontroller ](#a-microcontroller)
    - [a. issue](#a-issue)
    - [b. hardware programmer](#b-hardware-programmer)
  - [B. TinyGo coding](#b-tinygo-coding)
     - [a. Blinking LED](#a-blinking-led)  
     - [b. Install TinyGo](#b-install-tinygo)
     - [c. LED Fuction](#c-led-fuction)
     - [d. Scheduler](#d-scheduler)
     - [e. Test](#e-test)
    
## 1. Introduction

### A. Overview
SignAll, a manifacturing company of luminous signage needs a smart device for their product.
The Goal is to propose a devise with some customisation, setting and smart fuctionality for offer a better control of luminous signal.
In this technical, It has every technology, technical information needed for creating of device.

### B. Glossary
|Terms|Definition|
|-----|----------|
|Flashing|Writing a program to a microcontroller|

## 2. Development tools

### A. Software used
#### a. Technologies used
 | What ? | Which ? | Why ? |
 | :---------- | :----------: |----------: |
 | Code editor| [Visual Studio Code](https://code.visualstudio.com/)|/|
 |IDE|[IDE Arduino](https://www.arduino.cc/en/software)|Code editing for microcontroller device|
 | Repository| [GitHub](https://github.com/)|Code|
 | Managenment | [Trello](https://trello.com/)|Project Managing|
 |/|[OpenOCD](https://openocd.org/)|Board flashing|
 
 #### b. Language used 
 ![alt text](https://w3soft.org/wpub/media-pbld/2/l/langs-short-desc-go/go-logo.svg)
 
 [Golang](https://go.dev/)
 
 ver 1.19.1
 
 Go is a statically typed, compiled programming language designed at Google.
 It can be bring on microtroller via TinyGO.
 
 ![alt text](https://avatars.githubusercontent.com/u/45223613?s=280&v=4)
 
 [TinyGo](https://tinygo.org/)
 
 TinyGo is a Go compiler intended for use in small places such as microcontrollers, WebAssembly (Wasm), and command-line tools.
It can  run and compile TinyGo programs on [over  85 different microcontroller boards](https://tinygo.org/docs/reference/microcontrollers/).
 
 
 
 
 


### B. Hardware used
![alt text](https://seeklogo.com/images/S/seeed-studio-logo-4F3B000EB9-seeklogo.com.png)

 [Microcontroller LoRa-E5](https://tinygo.org/docs/reference/microcontrollers/lorae5/)
 
 HARDWARE SPECIFICATION
 
 ![image](https://user-images.githubusercontent.com/71770514/196367752-9721d48f-229c-4ca1-be54-d70ccb7a9069.png)

  PINOUT 
  
 ![image](https://user-images.githubusercontent.com/71770514/196367831-393646a0-9032-40a6-b874-549c14949e04.png)
 
[PIN](https://tinygo.org/docs/reference/microcontrollers/lorae5/#pins)
 
 ![alt text](https://cdn.iconscout.com/icon/free/png-256/arduino-1-226076.png)
 
 [STM32F103xx "bluepill" ](https://www.st.com/en/microcontrollers-microprocessors/stm32f103c8.html#overview)
 
 PINOUT
 
 ![image](https://www.electrokit.com/uploads/productfile/41016/stm32f103c8t6_pinout_voltage01.png)
 
 [PIN](https://tinygo.org/docs/reference/microcontrollers/bluepill/#pins)

## 3. Developement Process

### A. microcontroller

#### a. issue
Initially, for connect the microcontroller (Lora-E5) to the labtop or mac device and applie TinyGo code, we need to have any kind of device like USBDevice. However,  [Lora-E5 can't be compatible with USBDevice](https://tinygo.org/docs/reference/microcontrollers/lorae5/#interfaces), we need an other device. OpenOCD It's a software able to use a TinyGo for flashing Lora-E5 board and [ST-link](https://www.st.com/resource/en/user_manual/um1075-stlinkv2-incircuit-debuggerprogrammer-for-stm8-and-stm32-stmicroelectronics.pdf) device can be connect and use [TinyGo code on the board with SWD/SWIM headers](https://tinygo.org/docs/reference/microcontrollers/lorae5/#flashing). But we can't the board in this way and we need to find something else. We have too a stm32f1xx defice, we try with the IDE arduino to recognize the board but the IDE can't detect the device even we have installed the package and the driver.

#### b. hardware programmer
[STLinkV2](https://www.st.com/en/development-tools/st-link-v2.html) and [STLink](https://www.st.com/resource/en/user_manual/um1075-stlinkv2-incircuit-debuggerprogrammer-for-stm8-and-stm32-stmicroelectronics.pdf) it's a programmer needed for flashing.

STM32"bluepill" needs this followning step:
- Plug your STLink v2 programmer into your computer’s USB port.
- Plug your Bluepill into the STLink v2 programmer using the Bluepill [SWIO, SWCLK, 3V3 and GND pins](https://www.electrokit.com/uploads/productfile/41016/stm32f103c8t6_pinout_voltage01.png).
- Build and flash your TinyGo program using `tinygo flash -target=bluepill`

LoRa-E5 needs this following step:
- Connect an ST-Link device to the [SWD/SWIM headers](https://user-images.githubusercontent.com/71770514/196367831-393646a0-9032-40a6-b874-549c14949e04.png).
- Plug your LoRa-E5 board into your computer’s USB port.
- Plug your ST-Link device into your computer’s USB port.
- Build and flash your TinyGo program using `tinygo flash -target=lorae5`
### B. TinyGo coding

#### a. Blinking LED
[Blinking LED](https://tinygo.org/docs/tutorials/blinky/) it's the equivalent of `Hello World` of hardware, mean it's working.
Like any Golang code, it must to initialize a new Go module like this:
``` 
go mod init <name>

```
and creat blinking LED program:
``` go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 500)

        led.High()
        time.Sleep(time.Millisecond * 500)
    }
}
```
To transfer this program to the board, we use the `flash` subcommand. It's called  [Flashing](#b-glossary).
```
tinygo flash -target=circuitplay-express
```
`-target` flag depend on what board do you use, for exemple, if you used Arduino Uno, `-target=arduino` is necessary or `-target=itsybitsy-m4` for Adafruit ItsyBitsy M4.

`"machine"` import which provides direct access to the hardware in a somewhat portable way.

``` go
led.Configure(machine.PinConfig{Mode: machine.PinOutput})
```
This code confugure the GPIO pin, it need initialization before they can be used.

``` go
for {
    // ...
}
```
`for` here mean `while` in other languages (Golang don't have `while` statement). It's basicly an infinite loop in this case.

``` go
led.Low()
```
Set basicly 0V on the LED.

``` go
time.Sleep(time.Millisecond * 500)
```
It's here for delay the signal.

``` go
led.High()
```
Turn ON the LED.

#### b. Install TinyGo
On [MacOS](https://tinygo.org/getting-started/install/macos/), you need to install [Homebrew](https://brew.sh/)  for download tinygo.

If you don't have already install homebrew, type on your terminal:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```


and install tinygo:

```
    brew tap tinygo-org/tools 
    brew install tinygo
```

On [Windows](https://tinygo.org/getting-started/install/windows/), you need to install [Scoop](https://scoop.sh/) for download tinygo.

For install scoop, type on powershell:


```
irm get.scoop.sh | iex
```





however, if you get a message error like this ` PowerShell requires an execution policy in [Unrestricted, RemoteSigned, ByPass] to run Scoop. For example, to set the execution policy to 'RemoteSigned' please run 'Set-ExecutionPolicy RemoteSigned -Scope CurrentUser'.` you need to type before:
 
```
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time 
```




now, you can install scoop,

and install tinygo:

```
scoop install tinygo
```




#### c. LED Fuction

We have some feature with LED fuctions.

At first, we define the LED by a struct statement, and setup the LED.
``` go
type LED struct {
	pin machine.Pin
	on  bool
} 
```
 `pin` indicate what PIN on the board is using on the program.
 
 `on` basic setting ON by default of LED.
 
 After that, we need to set the pin on the board.
 ``` go
 leds := machine.PA12
	leds.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
 ```
 We use the STR32"bluepill" board, `PA12` is one of PIN of this board and we use it for the LED.
 
 We configure GPIO pin before using it with `Configure` function.
 
 after that, we applie the `leds` variable on LED struct
 
 ``` go
 l := LED{
		pin: leds,
		on:  false,
	}
 ```
 Now, we can use some function on the program.
 
 `On()` function is here for indicate if the LED is ON.
 
 `Set(value bool)` function set the LED on ON or OFF depend the argument is using.
 
`Blink()` function blink during a certain amont of time.

#### d. Scheduler

It's a function who you can set an hour when the LED is turn ON or OFF.

`func state(bool)` show if the LED is Turn ON or OFF.

`func scheduler()` set a hour you want to turn ON or OFF the LED.

`func scheduledOperation(hourOff int, minuteOff int, hourOn int, minuteOn int)` compare the current hour to setting hour and turn ON or OFF the LED depend of the hour.

#### e. Test

TDD exist too in Golang. For creat a new test file, you need to write `<name>_test.go`.

``` go
import "testing"

```
It's an import for use test function. For run your test file, you use the commande palette and type  `Go: Test File`.

example of test function:
``` go
func TestOn(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		want := true
		got := l.On()
		if got != want {
			t.Errorf("The light isn't turn on")
		}
	})
}
```
Basically, we testing our function if it get an expected result (here true bool). If the result is false, the test function return an error message and the test fail.

 
 
 
 






