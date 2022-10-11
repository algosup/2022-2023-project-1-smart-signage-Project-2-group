# Technical Specification
 2022-2023 Project Smart Signage Groupe 2
 
  Team:
  
- [Rémy Charles](https://github.com/RemyCHARLES)
- [Pierre Gorin](https://github.com/Pierre2103)
- [Élise Gautier](https://github.com/elisegtr)
- [Salaheddine Namir](https://github.com/T3rryc)
- [Grégory Pagnoux](https://github.com/Gregory-Pagnoux)

Create on Sep 27, 2022

Last Update Oct 11, 2022
## Table of contents


- [1. Introduction](#1-introduction)
- [2. Development tools](#2-development-tools)
- [3. Developement Process](#3-developement-process)

## 1. Introduction

### A. Overview
SignAll, a manifacturing company of luminous signage need a smart device for their product.
The Goal is to propose a devise with some customisation, setting and smart fuctionality for offer a better control of luminous signal.
In this technical, It has every technology, technical information needed for creating of device.

### B. Glossary
|Terms|Definition|
|-----|----------|
|Flashing|Writing a program to a microcontroller|

## 2. Development tools

### A. Software used
#### a.Technologies used
 | What ? | Which ? | Why ? |
 | :---------- | :----------: |----------: |
 | Code editor| [Visual Studio Code](https://code.visualstudio.com/)|/|
 |IDE|[IDE Arduino](https://www.arduino.cc/en/software)|Code editing for microcontroller device|
 | Repository| [GitHub](https://github.com/)|Code|
 | Managenment | [Trello](https://trello.com/)|Project Managing|
 |/|[OpenOCD](https://openocd.org/)|Board flashing|
 
 #### b.Language used
 ![alt text](https://w3soft.org/wpub/media-pbld/2/l/langs-short-desc-go/go-logo.svg)
 
 [Golang](https://go.dev/)
 
 ver 1.19.1
 
 Go is a statically typed, compiled programming language designed at Google.
 It can be bring on microtroller via TinyGO.
 
 ![alt text](https://avatars.githubusercontent.com/u/45223613?s=280&v=4)
 
 [TinyGo](https://tinygo.org/)
 
 TinyGo is a Go compiler intended for use in small places such as microcontrollers, WebAssembly (Wasm), and command-line tools.
It can  run and compile TinyGo programs on [over  85 different microcontroller boards](https://tinygo.org/docs/reference/microcontrollers/).
 
 
 
 
 


### B.Hardware used
![alt text](https://seeklogo.com/images/S/seeed-studio-logo-4F3B000EB9-seeklogo.com.png)

 [Microcontroller LoRa-E5](https://tinygo.org/docs/reference/microcontrollers/lorae5/)
 
 ![alt text](https://cdn.iconscout.com/icon/free/png-256/arduino-1-226076.png)
 
 [STM32F103xx ](https://www.st.com/en/microcontrollers-microprocessors/stm32f103c8.html#overview)

## 3. Developement Process

### A. microcontroller
Initially, for connect the microcontroller (Lora-E5) to the labtop or mac device and applie TinyGo code, we need to have any kind of device like USBDevice. However,  [Lora-E5 can't be compatible with USBDevice](https://tinygo.org/docs/reference/microcontrollers/lorae5/#interfaces), we need an other device. OpenOCD It's a software able to use a TinyGo for flashing Lora-E5 board and [ST-link](https://www.st.com/resource/en/user_manual/um1075-stlinkv2-incircuit-debuggerprogrammer-for-stm8-and-stm32-stmicroelectronics.pdf) device can be connect and use [TinyGo code on the board with SWD/SWIM headers](https://tinygo.org/docs/reference/microcontrollers/lorae5/#flashing). But we can't the board in this way and we need to find something else. We have too a stm32f1xx defice, we try with the IDE arduino to recognize the board but the IDE can't detect the device even we have installed the package and the driver.

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

#### b. Using PWM

#### c. GBD




