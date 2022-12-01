package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func run() error {
	readFile, err := os.Open("input.txt")

	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

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

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			updateMaxSums()
			currentSum = 0
			continue
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		currentSum += number
	}
	updateMaxSums()

	fmt.Printf("Part 1 result: %v\n", maxSums[0])

	var totalSum int
	for _, maxSum := range maxSums {
		totalSum += maxSum
	}
	fmt.Printf("Part 2 result: %v\n", totalSum)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
