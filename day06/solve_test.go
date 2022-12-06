package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, Solve, []advent.Test[[]int]{
		{
			File:     "input_sample.txt",
			Expected: []int{10, 29},
		},
		{
			File:     "input.txt",
			Expected: []int{1598, 2414},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, Solve, []string{"input.txt"})
}
