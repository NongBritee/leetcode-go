package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{input: 123, output: 321},
		{input: -123, output: -321},
		{input: 120, output: 21},
		{input: 0, output: 0},
		{input: 1534236469, output: 0},
		{input: -2147483648, output: 0},
		{input: 2147483647, output: 0},
	}

	for _, c := range cases {
		x := reverse(c.input)
		if x != c.output {
			t.Errorf("reverse(%d)=%d, expected %d", c.input, x, c.output)
		}
	}
}
