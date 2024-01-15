package leetcode

import "testing"

func TestValidSudoku(t *testing.T) {
	type testcase struct {
		name  string
		board [][]byte
		want  bool
	}

	tests := []testcase{
		{
			name: "Example1",
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: true,
		},
		{
			name: "Example2",
			board: [][]byte{
				[]byte("83..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := isValidSudoku(test.board)
			if ans != test.want {
				t.Errorf("got %t, want %t", ans, test.want)
			}
		})
	}
}
