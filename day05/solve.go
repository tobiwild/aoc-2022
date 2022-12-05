package puzzle

import (
	"bufio"
	"fmt"
	"io"
)

func Solve(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanningCargoLines := true

	var stacks1 [][]string
	var stacks2 [][]string

	updateStacks := func(stacks [][]string, line string) [][]string {
		for i, c := range line {
			if c >= 'A' && c <= 'Z' {
				if stacks == nil {
					stacks = make([][]string, (len(line)+1)/4)
				}
				si := i / 4
				stacks[si] = append([]string{string(c)}, stacks[si]...)
				continue
			}
		}
		return stacks
	}

	for scanner.Scan() {
		line := scanner.Text()

		if scanningCargoLines {
			if line == "" {
				scanningCargoLines = false
				continue
			}
			stacks1 = updateStacks(stacks1, line)
			stacks2 = updateStacks(stacks2, line)
			continue
		}

		var count, from, to int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			return nil, err
		}
		from--
		to--

		for i := 0; i < count; i++ {
			idx := len(stacks1[from]) - 1
			stacks1[to] = append(stacks1[to], stacks1[from][idx])
			stacks1[from] = stacks1[from][:idx]
		}

		{
			idx := len(stacks2[from]) - count
			stacks2[to] = append(stacks2[to], stacks2[from][idx:]...)
			stacks2[from] = stacks2[from][:idx]
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
