package main

import (
	"testing"
)

func Test_scheduler(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler()
		})
	}
}

func Test_scheduledOperation(t *testing.T) {
	type args struct {
		hourOff   int
		minuteOff int
		hourOn    int
		minuteOn  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduledOperation(tt.args.hourOff, tt.args.minuteOff, tt.args.hourOn, tt.args.minuteOn)
		})
	}
}
