package day02

import (
	"bufio"
	"fmt"
	"io"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	rock := 1
	paper := 2
	scissors := 3
	draw := 3
	win := 6

	isDraw := func(v1, v2 int) bool {
		return v1 == v2
	}

	rules := [][2]int{
		{rock, scissors},
		{scissors, paper},
		{paper, rock},
	}

	isWin := func(v1, v2 int) bool {
		for _, rule := range rules {
			win, losing := rule[0], rule[1]
			if v1 == win && v2 == losing {
				return true
			}
		}
		return false
	}

	getWinning := func(v int) int {
		for _, rule := range rules {
			win, losing := rule[0], rule[1]
			if v == losing {
				return win
			}
		}
		panic("invalid")
	}

	getLosing := func(v int) int {
		for _, rule := range rules {
			win, losing := rule[0], rule[1]
			if v == win {
				return losing
			}
		}
		panic("invalid")
	}

	p1 := map[string]int{"A": rock, "B": paper, "C": scissors}
	p2 := map[string]int{"X": rock, "Y": paper, "Z": scissors}

	getScore := func(v1, v2 int) int {
		score := v1
		if isDraw(v1, v2) {
			score += draw
		} else if isWin(v1, v2) {
			score += win
		}
		return score
	}

	scoreMap1 := make(map[string]int, 9)
	for alias1, v1 := range p1 {
		for alias2, v2 := range p2 {
			key := fmt.Sprintf("%v %v", alias1, alias2)
			// we only care for score of p2
			scoreMap1[key] = getScore(v2, v1)
		}
	}

	scoreMap2 := make(map[string]int, 9)
	for alias1, v1 := range p1 {
		for alias2 := range p2 {
			var v2 int
			switch alias2 {
			case "X":
				v2 = getLosing(v1)
			case "Y":
				v2 = v1
			case "Z":
				v2 = getWinning(v1)
			}
			key := fmt.Sprintf("%v %v", alias1, alias2)
			scoreMap2[key] = getScore(v2, v1)
		}
	}

	var totalScore1 int
	var totalScore2 int
	for scanner.Scan() {
		line := scanner.Text()

		if score, ok := scoreMap1[line]; ok {
			totalScore1 += score
		} else {
			return nil, fmt.Errorf("Invalid input: %v", line)
		}

		if score, ok := scoreMap2[line]; ok {
			totalScore2 += score
		} else {
			return nil, fmt.Errorf("Invalid input: %v", line)
		}
	}

	return []int{totalScore1, totalScore2}, nil
}
