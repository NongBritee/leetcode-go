package leetcode

import "testing"

func TestMinSteps(t *testing.T) {
	var tests = []struct {
		s, t string
		want int
	}{
		{"bab", "aba", 1},
		{"leetcode", "practice", 5},
		{"anagram", "mangaar", 0},
		{"xxyyzz", "xxyyzz", 0},
		{"friend", "family", 4},
	}

	for _, tt := range tests {
		got := minSteps(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("failed for %s, %s, expected %d, got %d", tt.s, tt.t, tt.want, got)
		}
	}
}
