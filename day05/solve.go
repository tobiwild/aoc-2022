package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var numberRegexp = regexp.MustCompile(`\d+`)

func Solve(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var cargoLines []string
	scanningCargoLines := true

	var stacks1 [][]string
	var stacks2 [][]string

	createStacks := func(stacks [][]string) [][]string {
		for i := len(cargoLines) - 2; i >= 0; i-- {
			for ci, c := range cargoLines[i] {
				stackIndex := ((ci + 3) / 4) - 1
				if stackIndex > len(stacks)-1 {
					stacks = append(stacks, []string{})
				}
				if c >= 'A' && c <= 'Z' {
					stacks[stackIndex] = append(stacks[stackIndex], string(c))
				}
			}
		}
		return stacks
	}

	for scanner.Scan() {
		line := scanner.Text()

		if scanningCargoLines {
			if line == "" {
				stacks1 = createStacks(stacks1)
				stacks2 = createStacks(stacks2)
				scanningCargoLines = false
				continue
			}
			cargoLines = append(cargoLines, line)
			continue
		}

		var nums []int
		numStrings := numberRegexp.FindAllString(line, 3)
		for _, numString := range numStrings {
			num, err := strconv.Atoi(string(numString))
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		if len(nums) != 3 {
			return nil, fmt.Errorf("Invalid line: %v", line)
		}
		count, from, to := nums[0], nums[1]-1, nums[2]-1

		for i := 0; i < count; i++ {
			flen := len(stacks1[from])
			top := stacks1[from][flen-1]
			stacks1[from] = stacks1[from][:flen-1]
			stacks1[to] = append(stacks1[to], top)
		}

		{
			flen := len(stacks2[from]) - count
			stacks2[to] = append(stacks2[to], stacks2[from][flen:]...)
			stacks2[from] = stacks2[from][:flen]
		}
	}

	getTopRow := func(stacks [][]string) string {
		var res []byte
		for _, stack := range stacks {
			slen := len(stack)
			if slen == 0 {
				continue
			}
			res = append(res, stack[slen-1]...)
		}
		return string(res)
	}

	return []string{getTopRow(stacks1), getTopRow(stacks2)}, nil
}
