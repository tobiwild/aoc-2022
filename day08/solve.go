package puzzle

import (
	"bufio"
	"io"
	"strconv"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var m [][]int

	for scanner.Scan() {
		line := scanner.Bytes()

		row := make([]int, 0, len(line))

		for _, numStr := range line {
			num, err := strconv.Atoi(string(numStr))
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}

		m = append(m, row)
	}

	n := len(m)
	visibleSum := 4 * (n - 1)
	var maxScenicStore int

	for y := 1; y < n-1; y++ {
		for x := 1; x < n-1; x++ {
			h := m[y][x]
			var blocked int
			var v1, v2, v3, v4 int
			for dx := x - 1; dx >= 0; dx-- {
				v1++
				if m[y][dx] >= h {
					blocked++
					break
				}
			}
			for dx := x + 1; dx < n; dx++ {
				v2++
				if m[y][dx] >= h {
					blocked++
					break
				}
			}
			for dy := y - 1; dy >= 0; dy-- {
				v3++
				if m[dy][x] >= h {
					blocked++
					break
				}
			}
			for dy := y + 1; dy < n; dy++ {
				v4++
				if m[dy][x] >= h {
					blocked++
					break
				}
			}
			scenicScore := v1 * v2 * v3 * v4
			if scenicScore > maxScenicStore {
				maxScenicStore = scenicScore
			}
			if blocked < 4 {
				visibleSum++
			}
		}
	}

	return []int{visibleSum, maxScenicStore}, nil
}
