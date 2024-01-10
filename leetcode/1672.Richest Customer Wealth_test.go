package leetcode

import "testing"

func TestMaximumWealth(t *testing.T) {

	type TestCase struct {
		name     string
		accounts [][]int
		want     int
	}

	cases := []TestCase{
		{"test 1", [][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{"test 2", [][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{"test 3", [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}

	for _, c := range cases {
		got := maximumWealth(c.accounts)
		if c.want != got {
			t.Errorf("accounts: %v\t|\tWant: %d\t|\tGot: %d", c.accounts, c.want, got)
		}
	}
}
