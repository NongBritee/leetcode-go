package leetcode

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	type test struct {
		name   string
		input  int
		expect int
	}

	tests := []test{
		{"test 1", 1, 1},
		{"test 2", 14, 6},
		{"test 3", 8, 4},
		{"test 4", 123, 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := numberOfSteps(test.input)
			if got != test.expect {
				t.Fatalf("numberOfSteps(%v) == %v, expect %v", test.input, got, test.expect)
			}
		})
	}
}
