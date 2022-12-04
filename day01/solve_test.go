package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, Solve, []advent.Test[[]int]{
		{
			File:     "input.txt",
			Expected: []int{66186, 196804},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, Solve, []string{"input.txt"})
}
