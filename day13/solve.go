package puzzle

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func nextItem(packet []byte) ([]byte, []byte) {
	var depth int

	for i, b := range packet {
		if i == 0 {
			if b != '[' {
				return packet, nil
			}
			continue
		}
		if b == '[' {
			depth++
		} else if b == ']' {
			depth--
		}
		if depth == 0 && b == ',' {
			return packet[1:i], append([]byte{'['}, packet[i+1:]...)
		}
		if i == len(packet)-2 {
			return packet[1 : i+1], nil
		}
	}

	return nil, nil
}

func compare(packet1, packet2 []byte) int {
	item1, rest1 := nextItem(packet1)
	item2, rest2 := nextItem(packet2)

	if len(item1) == 0 && len(item2) == 0 {
		return 0
	}
	if len(item1) == 0 {
		return -1
	}
	if len(item2) == 0 {
		return 1
	}

	if item1[0] == '[' || item2[0] == '[' {
		if c := compare(item1, item2); c != 0 {
			return c
		}
		return compare(rest1, rest2)
	}

	i1, err := strconv.Atoi(string(item1))
	if err != nil {
		panic(err)
	}
	i2, err := strconv.Atoi(string(item2))
	if err != nil {
		panic(err)
	}

	if i1 < i2 {
		return -1
	}

	if i1 > i2 {
		return 1
	}

	return compare(rest1, rest2)
}

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var index, result1 int
	packets := []string{"[[2]]", "[[6]]"}
	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}

		index++
		packet1 := sc.Text()
		sc.Scan()
		packet2 := sc.Text()

		if c := compare([]byte(packet1), []byte(packet2)); c == -1 {
			result1 += index
		}

		packets = append(packets, packet1, packet2)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare([]byte(packets[i]), []byte(packets[j])) == -1
	})

	result2 := 1
	for i, packet := range packets {
		if packet == "[[2]]" {
			result2 *= i + 1
			continue
		}

		if packet == "[[6]]" {
			result2 *= i + 1
			break
		}
	}

	return []int{result1, result2}, nil
}
