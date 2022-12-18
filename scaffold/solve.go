package puzzle

import (
	"bufio"
	"io"
)

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		// line := sc.Text()
	}

	return []int{}, nil
}
