package leetcode

import "testing"

func TestRotate(t *testing.T) {

	type testcase struct {
		name string
		nums []int
		k    int
		want []int
	}

	testcases := []testcase{
		{
			name: "test1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "test2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.nums, tc.k)
			for i := range tc.nums {
				if tc.nums[i] != tc.want[i] {
					t.Errorf("got: %v, want: %v", tc.nums, tc.want)
				}
			}
		})
	}
}
