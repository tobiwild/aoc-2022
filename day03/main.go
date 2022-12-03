package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func run() error {
	readFile, err := os.Open("input.txt")

	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	getPriority := func(c rune) int {
		if c <= 'Z' {
			return int(c-'A') + 27
		}
		return int(c-'a') + 1
	}

	var lineSum int
	var groupSum int
	groupMap := make(map[rune]int)

	for fileScanner.Scan() {
		line := fileScanner.Text()

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

	fmt.Printf("Part 1 result: %v\n", lineSum)

	fmt.Printf("Part 2 result: %v\n", groupSum)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
