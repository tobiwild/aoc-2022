package day01

import (
	"bufio"
	"io"
	"strconv"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var maxSums [3]int
	currentSum := 0

	updateMaxSums := func() {
		for i, maxSum := range maxSums {
			if currentSum > maxSum {
				maxSums[i] = currentSum
				break
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			updateMaxSums()
			currentSum = 0
			continue
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		currentSum += number
	}
	updateMaxSums()

	var totalSum int
	for _, maxSum := range maxSums {
		totalSum += maxSum
	}

	return []int{maxSums[0], totalSum}, nil
}
