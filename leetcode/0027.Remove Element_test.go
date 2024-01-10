package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	type testcase struct {
		name   string
		nums   []int
		val    int
		expect int
	}

	tests := []testcase{
		{
			name:   "test 1",
			nums:   []int{3, 2, 2, 3},
			val:    3,
			expect: 2,
		},
		{
			name:   "test 2",
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := removeElement(test.nums, test.val)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
