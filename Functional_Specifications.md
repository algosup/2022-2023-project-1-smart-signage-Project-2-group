# Functional Specification

## Table of contents
- [Functional Specification](#functional-specification)
  - [Table of contents](#table-of-contents)
    - [1. Distribution](#1-distribution)
    - [2. Overview](#2-overview)
    - [3. Scenarios](#3-scenarios)
        - [Scenario 1: Christophe](#scenario-1-christophe)
        - [Scenario 2: Arnaud](#scenario-2-arnaud)
    - [4. Risks and Assumptions](#4-risks-and-assumptions)
    - [5. Non goals](#5-non-goals)
    - [6. Development and environement and Requirements](#6-development-and-environement-and-requirements)
    - [7. Glossary](#7-glossary)

### 1. Distribution

| PERSON | ROLE |
| :-: | :-: |
| Remy CHARLES | Project Manager |
| Pierre GORIN | Software Engineer |
| Gregory PAGNOUX | Quality Assurance |
| Salahddine NAMIR | Tech Leader |
| Elise GAUTIER | Program Manager |
  
### 2. Overview

Smart signage requested by the signAll's[^1] company, is a service allowing to control the leds of the signs.

Actually, this company does not have any connected products. They are therefore obliged to move to find out if the signage is functional, on or off. In order to help them, we are going to create a product that will allow them to remotely control the luminous display of the signs but also to know the status of the LEDs[^4].

This spec is not, complete.

### 3. Scenarios

##### Scenario 1: Christophe

Christophe lives in Vierzon with his girlfriend. He is responsible for the signage of McDonald's fast food establishments at SignAll and comes by car. When fast food establishments close, he is responsible for turning off the signs. However, when the McDo restaurant closes, he is no longer in his office and must therefore do it from home. Thanks to the application, he can turn off all the McDonald's signs by sending a command message. Thanks to the application, he was able to turn off all the signs without having to go to his office.

##### Scenario 2: Arnaud

Arnaud lives in Chateauroux with his wife Caroline and these two children Angéline and Darius. Every morning, he goes to work by carpooling with his colleague Fabien. He is responsible for sign maintenance at SignAll and must do maintenance rounds each day to check on each panel if all the LEDs were working properly. It’s a very long and difficult task, but today his company has made a new update on their application that will help him accomplish his mission. On this application, it can simply perform maintenance tests without having to travel to each sign, thus reducing travel costs and saving time for the company.

### 4. Risks and Assumptions

**About the environment law :**

Signages are extinct between 1:00 a.m. and 6:00 a.m., when the signaled activity has ceased. If the activity begins or ceases between  0:00 a.m. and 7:00 a.m, signages must be turned off no later than one hour after the activity ends and may be turned on one hour before the activity resumes.

**the reglementation on the led panel :**

The surface density of installed luminous flux must not exceed 35 lm/ m² in built-up areas and 25 lm/m² outside built-up areas.

### 5. Non goals

This version will not support the following features:

  - An app or website who can allow us to see the status of the installation.

### 6. Development and environement and Requirements

  - Go[^2]
    - TinyGo[^3]
  - Windows/MacOS on development
  
### 7. Glossary

[^1]: SignAll
SignAll is a compagny that create innovative of the signages, customizable and as environmentally friendly as possible.

[^2]: GO
GO is an open source programming language used for general purpose. Go was developed by Google.

[^3]: TinyGO
TinyGo brings the Go programming language to embedded systems and to the modern web.

[^4]: LEDs
light-emitting diode (LED) is an electronic device that gives off light when it receives an electrical current.