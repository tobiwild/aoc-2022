package puzzle

import (
	"bufio"
	"io"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// line := scanner.Text()
	}

	return []int{}, nil
}
