# Quality Assurance

## Table of content

- [Quality Assurance](#quality-assurance)
	- [Table of content](#table-of-content)
	- [Test plan](#test-plan)
	- [History of bugs](#history-of-bugs)
				- [Test 1](#test-1)

## Test plan

*What is tested, how, and the expected and achieved results ?*

| REFERENCE TEST | ACTION | RESULT EXPECTED | RESULT ACHIEVED |
| :-: | :-: | :-: | :-: |
| 1a | Press the button to turn on the light | light on | **FAIL** |
| 1b |  |  | **FAIL** |
| 1c |  |  |  |
| 2a | Press the button to turn off the light | light off |  |
| 3a | Press the button to blink the light | light blink |  |
| 4a | Modify the lightness | precentage of light |  |
| 5a | put the light detector in the dark | the brightness of the light increases |  |
| 6a | put the light detector in the daylight | the brightness of the light decreases |  |
| 7a | at a given time, the lights turn on | light on |  |
| 8a | at a given time, the lights turn off | light off |  |
| 9a | turn on all the lights except one (voluntarily turned off) | displays an alert message |  |
| 10a | put voluntarily a high temperature | displays an alert message |  |

## History of bugs

*(refer to the table)*

##### Test 1

	type LED struct {
		pin machine.Pin
		on  bool
	}

	func New() *LED {
		leds := machine.PA12
		leds.Configure(machine.PinConfig{
			Mode: machine.PinOutput,
		})
		l := LED{
			pin: leds,
			on:  false,
		}
		return &l
	}

	func (l *LED) On() bool {
		return l.on
	}



	func TestOn(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			want := true
			got := l.On()
			if got != want {
				t.Errorf("The light isn't turn on")
			}
		})
	}

| REFERENCE TEST | TEST REALIZED | EXPLICATION OF THE BUG |
| :-: | :-: | :-: |
| a | when you enter the string `"true"`, the LED blinks and returns `true`. | the test didn't return an answer. |
| b | when you call the fonction `l.On()`, return an error if the LED is on  | the test return that `l` is undefined |
| c |  |  |
