# Technical Specification

## Table of contents


- [1. Introduction](#1-introduction)
- [2. Development tools](#2-development-tools)
- [3. Developement Process](#3-developement-process)

## 1. Introduction
SignAll, a manifacturing company of luminous signage need a smart device for their product.
The Goal is to propose a devise with some customisation, setting and smart fuctionality for offer a better control of luminous signal.
In this technical, It has every technology, technical information needed for creating of device.

## 2. Development tools

### A. Software used
#### a.Technologies used
 | What ? | Which ? | Why ? |
 | :---------- | :----------: |----------: |
 | Code editor| [Visual Studio Code](https://code.visualstudio.com/)|/|
 |IDE|[IDE Arduino](https://www.arduino.cc/en/software)|Code editing|
 | Repository| [GitHub](https://github.com/)|Code|
 | Managenment | [Trello](https://trello.com/)|Project Managing|
 |/|[OpenOCD](https://openocd.org/)|Board flashing|
 
 #### b.Language used
 ![alt text](https://w3soft.org/wpub/media-pbld/2/l/langs-short-desc-go/go-logo.svg)
 
 [Golang](https://go.dev/)
 
 ![alt text](https://avatars.githubusercontent.com/u/45223613?s=280&v=4)
 
 [TinyGo](https://tinygo.org/)
 
 
 


### B.Hardware used
![alt text](https://seeklogo.com/images/S/seeed-studio-logo-4F3B000EB9-seeklogo.com.png)

 [Microcontroller LoRa-E5](https://tinygo.org/docs/reference/microcontrollers/lorae5/)
 
 ![alt text](https://cdn.iconscout.com/icon/free/png-256/arduino-1-226076.png)
 
 [STM32F103xx ](https://www.st.com/en/microcontrollers-microprocessors/stm32f103c8.html#overview)

## 3. Developement Process

### A. microcontroller
Initially, for connect the microcontroller (Lora-E5) to the labtop or mac device and applie TinyGo code, we need to have any kind of device like USBDevice. However,  Lora-E5 can't be compatible with USBDevice, we need an other device. OpenOCD It's a software able to use a TinyGo for flashing Lora-E5 board and ST-link device can be connect and use TinyGo code on the board with SWD/SWIM headers. But we can't the board in this way and we need to find something else.    

