package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type testcase struct {
		name   string
		input  []int
		expect int
	}

	tests := []testcase{
		{
			name:   "Example 1",
			input:  []int{1, 1, 2},
			expect: 2,
		},
		{
			name:   "Example 2",
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := removeDuplicates(test.input)
			if got != test.expect {
				t.Fatalf("removeDuplicates(%v) = %v, but expect %v", test.input, got, test.expect)
			}
		})
	}
}
