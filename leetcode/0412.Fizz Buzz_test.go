package leetcode

import "testing"

func TestFizzBuzz(t *testing.T) {
	type testcase struct {
		name   string
		input  int
		expect []string
	}

	tests := []testcase{
		{"test 1", 1, []string{"1"}},
		{"test 2", 2, []string{"1", "2"}},
		{"test 3", 3, []string{"1", "2", "Fizz"}},
		{"test 4", 4, []string{"1", "2", "Fizz", "4"}},
		{"test 5", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"test 6", 6, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"}},
		{"test 7", 7, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}},
		{"test 8", 8, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8"}},
		{"test 9", 9, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz"}},
		{"test 10", 10, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}},
		{"test 11", 11, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11"}},
		{"test 12", 12, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz"}},
		{"test 13", 13, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13"}},
		{"test 14", 14, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14"}},
		{"test 15", 15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			got := fizzBuzz(c.input)
			if len(got) != len(c.expect) {
				t.Fatalf("fizzBuzz(%v) == %v, expect %v", c.input, got, c.expect)
			}
			for i := range got {
				if got[i] != c.expect[i] {
					t.Fatalf("fizzBuzz(%v) == %v, expect %v", c.input, got, c.expect)
				}
			}
		})
	}
}
