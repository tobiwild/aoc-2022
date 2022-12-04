package day04

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var includeCount int
	var overlapCount int
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			numberStrings := strings.Split(part, "-")
			for _, numberString := range numberStrings {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, number)
			}
		}
		if len(numbers) != 4 {
			return nil, fmt.Errorf("Invalid line: %v", line)
		}
		left := max(numbers[0], numbers[2])
		right := min(numbers[1], numbers[3])

		if right >= left {
			overlapCount++
			for i := 0; i < 4; i += 2 {
				if numbers[i] >= left && numbers[i+1] <= right {
					includeCount++
					break
				}
			}
		}
	}

	return []int{includeCount, overlapCount}, nil
}
