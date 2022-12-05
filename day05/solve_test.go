package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, Solve, []advent.Test[[]string]{
		{
			File:     "input_sample.txt",
			Expected: []string{"CMZ", "MCD"},
		},
		{
			File:     "input.txt",
			Expected: []string{"ZWHVFWQWW", "HZFZCCWWV"},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, Solve, []string{"input.txt"})
}
