# Functional Specification

## Table of contents
- [Functional Specification](#functional-specification)
  - [Table of contents](#table-of-contents)
    - [1. Distribution](#1-distribution)
    - [1. Overview](#1-overview)
    - [2. Scenarios](#2-scenarios)
        - [Scenario 1: Christophe](#scenario-1-christophe)
        - [Scenario 2: Arnaud](#scenario-2-arnaud)
    - [3. goals](#3-goals)
    - [4. Non goals](#4-non-goals)
    - [5. Development and environement and Requirements](#5-development-and-environement-and-requirements)
    - [6. Glossary](#6-glossary)

### 1. Distribution

| PERSON | ROLE |
| :-: | :-: |
| Remy CHARLES | Poject Manager |
| Pierre GORIN | Software Engineer |
| Gregory PAGNOUX | Quality Assurance |
| Salahddine NAMIR | Tech Leader |
| Elise GAUTIER | Program Manager |
  
### 1. Overview

Smart signage requested by the signAll's[^1] company, is a service allowing to control the leds of the signs.

Actually, this company does not have any connected products. They are therefore obliged to move to find out if the signage is functional, on or off. In order to help them, we are going to create a product that will allow them to remotely control the luminous display of the signs but also to know the status of the LEDs[^4].

This spec is not, by any stretch of the imagination, complete.

### 2. Scenarios

##### Scenario 1: Christophe

He is responsible for the signage of fast food establishments (McDo, Burger King, etc.) at SignAll. During the closure of fast food establishments, he is responsible for turning off and on the signs respectively from 1:00 a.m. to 6:00 a.m. in order to comply with regulations. Using the application, he will be able to carry out his mission for all the fast food restaurants in France without having to go to each fast food restaurant.

##### Scenario 2: Arnaud

He is responsible for sign maintenance at SignAll. His job was to do maintenance rounds to check on each panel if all the LEDs were working properly. Shortly afterwards, thanks to the application, he can simply carry out maintenance tests remotely and thus allows the company to reduce his travel costs.

### 3. goals

This version will support the following features:

  - Activation and deactivation of remote signaling.
  - Remote monitoring of signaling status (functional or faulty).
  - Brightness change at fixed time.
  - Adaptation of the intensity of the signage according to the external luminosity.
  - Real-time fault location.
  - Create a watchdog to know if the signage can be repaired remotely or if it requires the intervention of a technician.
  - Leave the choice to the company to adjust the start time and the end time of ignition according to its needs.
  - The ability to update the entire panel at once.
  - Warn about some over-heat.
  - Send an alert to know when the panel needs to be cleaned or checked.

### 4. Non goals

This version will not support the following features:

  - An app or website who can allow us to see the status of the installation.

### 5. Development and environement and Requirements

  - Go[^2]
    - TinyGo[^3]
  - Windows/MacOS on development
  
### 6. Glossary

[^1]: signAll
SignAll is a compagny that create innovative of the signages, customizable and as environmentally friendly as possible.

[^2]: GO
GO is an open source programming language used for general purpose. Go was developed by Google.

[^3]: TinyGO
TinyGo brings the Go programming language to embedded systems and to the modern web.

[^4]: LEDs
light-emitting diode (LED) is an electronic device that gives off light when it receives an electrical current.