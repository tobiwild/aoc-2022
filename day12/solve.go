package puzzle

import (
	"bufio"
	"io"
)

type pos struct {
	x, y int
}

type cell struct {
	char  byte
	steps int
}

func (c cell) height() byte {
	if c.char == 'S' {
		return 'a'
	}
	if c.char == 'E' {
		return 'z'
	}
	return c.char
}

type grid [][]*cell

func (g grid) reset() {
	for _, row := range g {
		for _, cell := range row {
			cell.steps = 0
		}
	}
}

var dirs = []pos{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type stepper struct {
	grid grid
	stop func(cell *cell) bool
	down bool
}

func (s stepper) step(p pos, steps int) int {
	g := s.grid
	cell := g[p.y][p.x]
	if cell.steps > 0 && steps >= cell.steps {
		return -1
	}
	cell.steps = steps

	if s.stop(cell) {
		return steps
	}

	ly := len(g)
	lx := len(g[p.y])
	result := -1
	for _, d := range dirs {
		n := pos{y: p.y + d.y, x: p.x + d.x}
		if n.y < 0 || n.y >= ly || n.x < 0 || n.x >= lx {
			continue
		}
		newCell := g[n.y][n.x]
		if s.down {
			if int(newCell.height())-int(cell.height()) < -1 {
				continue
			}
		} else {
			if int(newCell.height())-int(cell.height()) > 1 {
				continue
			}
		}
		stepResult := s.step(n, steps+1)
		if stepResult > -1 && (result == -1 || stepResult < result) {
			result = stepResult
		}
	}
	return result
}

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var grid grid
	var start, end pos

	var y int
	for sc.Scan() {
		row := make([]*cell, len(sc.Bytes()))
		for x, char := range sc.Bytes() {
			if char == 'S' {
				start = pos{y: y, x: x}
			}
			if char == 'E' {
				end = pos{y: y, x: x}
			}
			row[x] = &cell{char: char}
		}

		grid = append(grid, row)
		y++
	}

	steps1 := (stepper{
		grid: grid,
		stop: func(cell *cell) bool { return cell.char == 'E' },
	}).step(start, 0)

	grid.reset()
	steps2 := (stepper{
		grid: grid,
		stop: func(cell *cell) bool { return cell.height() == 'a' },
		down: true,
	}).step(end, 0)

	return []int{steps1, steps2}, nil
}
