package puzzle

import (
	"bufio"
	"io"
	"unicode/utf8"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	getPriority := func(c rune) int {
		if c <= 'Z' {
			return int(c-'A') + 27
		}
		return int(c-'a') + 1
	}

	var lineSum int
	var groupSum int
	groupMap := make(map[rune]int)

	for scanner.Scan() {
		line := scanner.Text()

		{
			cMap := make(map[rune]struct{})
			midIndex := utf8.RuneCountInString(line) / 2
			for i, c := range line {
				if i >= midIndex {
					if _, ok := cMap[c]; ok {
						lineSum += getPriority(c)
						break
					}
					continue
				}
				cMap[c] = struct{}{}
			}
		}

		{
			cMap := make(map[rune]struct{})
			for _, c := range line {
				if _, ok := cMap[c]; ok {
					continue
				}
				cMap[c] = struct{}{}
				groupMap[c]++
				if groupMap[c] == 3 {
					groupSum += getPriority(c)
					groupMap = make(map[rune]int)
					break
				}
			}
		}
	}

	return []int{lineSum, groupSum}, nil
}
