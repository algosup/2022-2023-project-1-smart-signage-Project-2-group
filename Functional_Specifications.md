# Functional Specification

## Document tracking

| Vers. | Date | Author | Subject of the update |
| :-: | :-: | :-: | :-: |
| 1.0 | 09/27/2022 | Elise Gautier | Document creation |
| 2.0 | 09/29/2022 | Elise Gautier | Insertion of information as part of the project |
| 3.0 | 10/07/2022 | Elise Gautier | Inserting the distribution and risks and assumptions|
| 4.0 | 10/13/2022 | Elise Gautier | Inserting objective of the project |
| 5.0 | 10/21/2022 | Elise Gautier | Inserting personas |

## Table of contents

- [Functional Specification](#functional-specification)
  - [Document tracking](#document-tracking)
  - [Table of contents](#table-of-contents)
    - [1. Distribution](#1-distribution)
    - [2. Overview](#2-overview)
    - [3. Personas](#3-personas)
        - [Christophe](#christophe)
        - [Arnaud](#arnaud)
        - [Alix](#alix)
        - [Maxence](#maxence)
      - [3.1 Scenarios](#31-scenarios)
        - [Scenario 1: Christophe](#scenario-1-christophe)
        - [Scenario 2: Arnaud](#scenario-2-arnaud)
        - [Scenario 3: Alix](#scenario-3-alix)
    - [4. objective of the project](#4-objective-of-the-project)
    - [5. Risks and Assumptions](#5-risks-and-assumptions)
    - [6. Non goal](#6-non-goal)
    - [7. Development and environement and Requirements](#7-development-and-environement-and-requirements)
    - [8. Glossary](#8-glossary)

### 1. Distribution

| PERSON | ROLE |
| :-: | :-: |
| Remy CHARLES | Project Manager |
| Pierre GORIN | Software Engineer |
| Gregory PAGNOUX | Quality Assurance |
| Salaheddine NAMIR | Tech Leader |
| Elise GAUTIER | Program Manager |
  
### 2. Overview

Appsolu requested by the signAll's[^1] company, is a service allowing to control the leds of the signs.

Actually, this company does not have any connected products. They are therefore obliged to move to find out if the signage is functional, on or off. In order to help them, we are going to create a product that will allow them to remotely control the luminous display of the signs but also to know the status of the LEDs[^4].

### 3. Personas

##### Christophe

**Age:** 30
**Location:** Vierzon
**Marital Status:** Maried
**Kids:** 2 children
**Occupation:** Sector Manager
**Workplace:** SignAll
**Pets:** 1 dog
**means of transport:** He comes by car

**Characteristics:**

* Reliable
* Caring
* Committed
* Action Oriented
* Focused
* Eager

**Personal life:**
He has been working at SignAll for 3 years as a sector manager. He has since may the new SignAll application which will help him in his work because he is responsible for the signage.

##### Arnaud

**Age:** 38
**Location:** Vignoux-sur-Barangeon
**Marital Status:** Engaged with his wife Caroline
**Kids:** 2 children (Angéline and Darius)
**Occupation:** Responsible for sign maintenance
**Workplace:** SignAll
**means of transport:** He goes to work by carpooling with his colleague Fabien

**Characteristics:**

* Focused
* Reliable
* Organized
* Effective

**Personal life:**
Arnaud was part of a maintenance company but he came to SignAll last year. He must use his company car to carry out maintenance checks.

##### Alix

**Age:** 25
**Location:** Châteauroux
**Marital Status:** single
**Kids:** No
**Occupation:** employee
**Workplace:** AXA's bank
**Pets:** 1 kitten
**means of transport:** She goes to work in carpooling with his roommate because she doesn't the license

**Characteristics:**

* Committed
* Reliable
* Focused
* Organized

**Personal life:**
Alix has been working at AXA for about 3 years. She makes the journey every day with her roommate whom she has known for 5 years now. Alix is ​​a very sporty and dynamic person. She also recently adopted a kitten with her roommate.

##### Maxence

**Age:** 50
**Location:** Lyon
**Marital Status:** Divorced
**Kids:** 1 child
**Occupation:** SignAll CEO
**Workplace:** SignAll
**means of transport:** he comes by car

**Characteristics:**

* Eager
* Committed
* Focused
* Organized

**Personal life:**
Maxence has been with the company for 20 years and became CEO 5 years ago. He often goes to meetings in other cities or even in other countries but always manages to stay close to his employees. Thanks to the Absolut application acquired by the company, it will be able to continue to help its employees even abroad.

#### 3.1 Scenarios

##### Scenario 1: Christophe

Christophe is responsible for the signage of McDonald's.
When fast food establishments close, he is responsible for turning off the signs. However, when the McDo restaurant closes, he is no longer in his office and must therefore do it from home.
Thanks to the application, he can turn on all the McDonald's signs by sending a command message. Thanks to the application, he was able to turn off all the signs without having to go to his office.

##### Scenario 2: Arnaud

He must do maintenance rounds each day to check on each panel if all the LEDs were working properly. It's a very long and difficult task, but today his company has made a new update on their application that will help him accomplish his mission faster.
On this application, it can simply perform maintenance tests without having to travel to each sign, thus reducing travel costs and saving time for the company.

##### Scenario 3: Alix

Like every morning, she has to open the bank and lights off the sign. However, she don't have the qualification to do this from the app. She will therefore call her boss Maxence who is not in the same city as the bank because he is going to a budget meeting. Maxence can with a send command, remotely control lights on and off.
Thanks to this he will be able to send a command message so that the lights go out at the bank without having to move. Alix is able to continue his work and open the bank for the future clients like every other day.

### 4. objective of the project

For this product, we want to receive a message to see remotely the signage status in real time. The main features of the project are:

- A follow-up of the signage to know if it is on or off.
- Sent a notification if there is a product failure.
- Sent a notification if there is overheating of the product.
- A reduction in consumption such as stopping at a fixed time.
- Adjustment of the intensity according to the ambient lighting.
  
### 5. Risks and Assumptions

**About the environment law :**

Signages are extinct between 1:00 a.m. and 6:00 a.m., when the signaled activity has ceased. If the activity begins or ceases between  0:00 a.m. and 7:00 a.m, signages must be turned off no later than one hour after the activity ends and may be turned on one hour before the activity resumes.

**the reglementation on the led panel :**

The surface density of installed luminous flux must not exceed 35 lm/ m² in built-up areas and 25 lm/m² outside built-up areas.

### 6. Non goal

This version will not support the following features:

- An app or website who can allow us to see the status of the installation.

### 7. Development and environement and Requirements

- Go[^2]
  - TinyGo[^3]
- Windows/MacOS on development
  
### 8. Glossary

[^1]: SignAll
SignAll is a compagny that create innovative of the signages, customizable and as environmentally friendly as possible.

[^2]: GO
GO is an open source programming language used for general purpose. Go was developed by Google.

[^3]: TinyGO
TinyGo brings the Go programming language to embedded systems and to the modern web.

[^4]: LEDs
light-emitting diode (LED) is an electronic device that gives off light when it receives an electrical current.
