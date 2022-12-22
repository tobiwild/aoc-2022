package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, puzzle{part1Y: 10, part2Max: 20}.Solve, []advent.Test[[]int]{
		{
			File:     "input_sample.txt",
			Expected: []int{26, 56000011},
		},
	})
	advent.TestSolve(t, puzzle{part1Y: 2000000, part2Max: 4000000}.Solve, []advent.Test[[]int]{
		{
			File:     "input.txt",
			Expected: []int{5073496, 13081194638237},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, puzzle{part1Y: 2000000, part2Max: 4000000}.Solve, []string{"input.txt"})
}
