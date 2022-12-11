package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
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
	items        []int
	operation    operation
	test         test
	inspectCount int
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

// better error message
func sscanf(str string, format string, a ...any) error {
	_, err := fmt.Sscanf(str, format, a...)
	if err != nil {
		return fmt.Errorf("Line %q does not match %q: %w", str, format, err)
	}
	return nil
}

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var monkeys []monkey

	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}

		// Monkey 1:
		m := monkey{}

		// Starting items: 54, 65, 75, 74
		{
			sc.Scan()
			numStrings := numberRegexp.FindAllString(sc.Text(), -1)
			for _, numStr := range numStrings {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}
				m.items = append(m.items, num)
			}
		}

		// Operation: new = old + 6
		{
			sc.Scan()
			var operandStr, valueString string
			if err := sscanf(sc.Text(), " Operation: new = old %s %s", &operandStr, &valueString); err != nil {
				return nil, err
			}
			if value, err := strconv.Atoi(valueString); err == nil {
				m.operation.value = &value
			}
			var operand operand
			if operandStr == "+" {
				operand = OPERAND_PLUS
			} else if operandStr == "*" {
				operand = OPERAND_MULTIPLY
			} else {
				return nil, fmt.Errorf("Unknown operand %v", operandStr)
			}
			m.operation.operand = operand
		}

		// Test: divisible by 19
		{
			sc.Scan()
			if err := sscanf(sc.Text(), " Test: divisible by %d", &m.test.divisibleBy); err != nil {
				return nil, err
			}
		}

		// If true: throw to monkey 2
		{
			sc.Scan()
			if err := sscanf(sc.Text(), " If true: throw to monkey %d", &m.test.trueMonkey); err != nil {
				return nil, err
			}
		}

		// If false: throw to monkey 0
		{
			sc.Scan()
			if err := sscanf(sc.Text(), " If false: throw to monkey %d", &m.test.falseMonkey); err != nil {
				return nil, err
			}
		}

		monkeys = append(monkeys, m)
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
		for round := 0; round < rounds; round++ {
			for mi, monkey := range monkeys {
				for _, item := range monkey.items {
					monkeys[mi].inspectCount++
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
		inspections := make([]int, len(monkeys))
		for i, monkey := range monkeys {
			inspections[i] = monkey.inspectCount
		}
		sort.Ints(inspections)
		l := len(inspections)
		return inspections[l-2] * inspections[l-1]
	}

	monkeyBusiness1 := play(copyMonkeys(monkeys), 20, true)
	monkeyBusiness2 := play(copyMonkeys(monkeys), 10_000, false)

	return []int{monkeyBusiness1, monkeyBusiness2}, nil
}
