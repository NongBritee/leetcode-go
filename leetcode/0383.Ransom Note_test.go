package leetcode

import (
	"testing"
)

func TestRandomNote(t *testing.T) {
	type TestCase struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}

	cases := []TestCase{
		{"test 1", "a", "b", false},
		{"test 2", "aa", "ab", false},
		{"test 3", "aa", "aab", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans := canConstruct(c.ransomNote, c.magazine)
			if ans != c.want {
				t.Errorf("got %v, want %v", ans, c.want)
			}
		})
	}
}
