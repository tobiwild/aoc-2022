package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type operand int

const (
	OPERAND_MULTIPLY operand = iota
	OPERAND_PLUS
)

type operation struct {
	operand operand
	value   *int
}

type test struct {
	divisibleBy int
	trueMonkey  int
	falseMonkey int
}

type monkey struct {
	items     []int
	operation operation
	test      test
}

func copyMonkeys(monkeys []monkey) []monkey {
	result := make([]monkey, len(monkeys))
	for i, monkey := range monkeys {
		items := make([]int, len(monkey.items))
		copy(items, monkey.items)
		monkey.items = items
		result[i] = monkey
	}
	return result
}

var (
	numberRegexp = regexp.MustCompile(`\d+`)
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var monkeys []monkey

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, monkey{})
			continue
		}

		if len(monkeys) == 0 {
			return nil, fmt.Errorf("No active monkey")
		}

		curMonkey := &monkeys[len(monkeys)-1]

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Starting items:") {
			numStrings := numberRegexp.FindAllString(line, -1)
			for _, numStr := range numStrings {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}
				curMonkey.items = append(curMonkey.items, num)
			}
			continue
		}

		if strings.HasPrefix(line, "Operation:") {
			var operandStr, valueString string

			if _, err := fmt.Sscanf(line, "Operation: new = old %s %s", &operandStr, &valueString); err != nil {
				return nil, fmt.Errorf("Line %q invalid: %v", line, err)
			}
			if value, err := strconv.Atoi(valueString); err == nil {
				curMonkey.operation.value = &value
			}
			var operand operand
			if operandStr == "+" {
				operand = OPERAND_PLUS
			} else if operandStr == "*" {
				operand = OPERAND_MULTIPLY
			} else {
				return nil, fmt.Errorf("Unknown operand %v", operandStr)
			}
			curMonkey.operation.operand = operand
			continue
		}

		if strings.HasPrefix(line, "Test:") {
			var divisibleBy int
			if _, err := fmt.Sscanf(line, "Test: divisible by %d", &divisibleBy); err != nil {
				return nil, fmt.Errorf("Line %q invalid: %v", line, err)
			}
			curMonkey.test.divisibleBy = divisibleBy
			continue
		}

		if strings.HasPrefix(line, "If") {
			var boolValue string
			var toMonkey int
			if _, err := fmt.Sscanf(line, "If %s throw to monkey %d", &boolValue, &toMonkey); err != nil {
				return nil, fmt.Errorf("Line %q invalid: %v", line, err)
			}
			if boolValue == "true:" {
				curMonkey.test.trueMonkey = toMonkey
			} else {
				curMonkey.test.falseMonkey = toMonkey
			}
			continue
		}
	}

	// A = level
	// B = Monkey's "divisible by"
	//
	// see https://de.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/what-is-modular-arithmetic
	//
	// A mod B = (A + K * B) mod B
	// 3  mod 10 = 3
	// 13 mod 10 = 3
	// 23 mod 10 = 3
	// 33 mod 10 = 3
	//
	// So when we set K = product of all possible B's
	// we can use it for all monkeys (every monkey has its B included so we comply with the rule from above)
	// and do level %= commonProduct to reduce the level
	// see also https://en.wikipedia.org/wiki/Modular_arithmetic
	commonProduct := 1
	for _, monkey := range monkeys {
		commonProduct *= monkey.test.divisibleBy
	}

	play := func(monkeys []monkey, rounds int, divideLevel bool) int {
		inspections := make(map[int]int)
		for round := 0; round < rounds; round++ {
			for mi, monkey := range monkeys {
				for _, item := range monkey.items {
					inspections[mi]++
					level := item
					value := level
					if monkey.operation.value != nil {
						value = *monkey.operation.value
					}
					if monkey.operation.operand == OPERAND_MULTIPLY {
						level *= value
					} else {
						level += value
					}
					if divideLevel {
						level /= 3
					}
					level %= commonProduct
					if level%monkey.test.divisibleBy == 0 {
						t := monkey.test.trueMonkey
						monkeys[t].items = append(monkeys[t].items, level)
					} else {
						t := monkey.test.falseMonkey
						monkeys[t].items = append(monkeys[t].items, level)
					}
				}
				monkeys[mi].items = nil
			}
		}
		inspectionCounts := make([]int, len(inspections))
		for i, count := range inspections {
			inspectionCounts[i] = count
		}
		sort.Ints(inspectionCounts)
		l := len(inspectionCounts)
		return inspectionCounts[l-2] * inspectionCounts[l-1]
	}

	monkeyBusiness1 := play(copyMonkeys(monkeys), 20, true)
	monkeyBusiness2 := play(copyMonkeys(monkeys), 10_000, false)

	return []int{monkeyBusiness1, monkeyBusiness2}, nil
}
