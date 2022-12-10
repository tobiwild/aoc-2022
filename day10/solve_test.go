package puzzle

import (
	"testing"

	advent "github.com/tobiwild/aoc-2022"
)

func TestSolve(t *testing.T) {
	advent.TestSolve(t, Solve, []advent.Test[*result]{
		{
			File: "input_sample.txt",
			Expected: &result{
				signalStrength: 13140,
				screenImage: `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`,
			},
		},
		{
			File: "input.txt",
			Expected: &result{
				signalStrength: 14920,
				screenImage: `
###..#..#..##...##...##..###..#..#.####.
#..#.#..#.#..#.#..#.#..#.#..#.#..#....#.
###..#..#.#....#..#.#....###..#..#...#..
#..#.#..#.#....####.#....#..#.#..#..#...
#..#.#..#.#..#.#..#.#..#.#..#.#..#.#....
###...##...##..#..#..##..###...##..####.
`, // BUCACBUZ
			},
		},
	})
}

func BenchmarkSolve(b *testing.B) {
	advent.BenchmarkSolve(b, Solve, []string{"input.txt"})
}
