# Quality Assurance

### Table of content

- [Test plan](#test-plan)
- [history of bugs](#history-of-bugs)
  - [Test 1a](#Test-1a)
  - [Test 1b](#Test-1b)
  - [Test 1c](#Test-1c)

### Test plan

###### *What is tested, how, and the expected and achieved results ?*

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

### history of bugs

###### *(refer to the table)*

##### Test 1a

    func TestLight_on(t *testing.T) {
        t.Run("true", func(t *testing.T) {
		    val := "true"
		    got := light_on(val)
		    if got != "true" {
			    fmt.Println("The light isn't turn on")
		    }
	    })
    }
It's a fail because

##### Test 1b

    func TestOn(t *testing.T) {
	    t.Run("true", func(t *testing.T) {
		    want := true
		    got := l.On()
		    if got != want {
			    t.Errorf("The light isn't turn on")
		    }
	    })
    }
It's fail beacause

##### Test 1c
