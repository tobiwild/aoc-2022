package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, Solve, []advent.Test[[]int]{
		{
			File:     "input_sample.txt",
			Expected: []int{13, 1},
		},
		{
			File:     "input_sample2.txt",
			Expected: []int{88, 36},
		},
		{
			File:     "input.txt",
			Expected: []int{6642, 2765},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, Solve, []string{"input.txt"})
}
