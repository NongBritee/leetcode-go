package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {

	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "Example 1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "Example 2",
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name: "Example 3",
			s:    []byte{'a'},
			want: []byte{'a'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("got: %v, want: %v", tt.s, tt.want)
			}

		})
	}
}
