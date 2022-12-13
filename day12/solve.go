package puzzle

import (
	"bufio"
	"io"
)

type pos struct {
	x, y int
}

type cell struct {
	char    byte
	visited bool
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
			cell.visited = false
		}
	}
}

var dirs = []pos{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type queueItem struct {
	pos   pos
	steps int
}

func shortestPath(g grid, queue []queueItem) int {
	for len(queue) > 0 {
		q := queue[0]
		p := q.pos
		queue = queue[1:]
		cell := g[p.y][p.x]

		if cell.visited {
			continue
		}
		cell.visited = true

		if cell.char == 'E' {
			return q.steps
		}

		ly := len(g)
		lx := len(g[p.y])
		for _, d := range dirs {
			n := pos{y: p.y + d.y, x: p.x + d.x}
			if n.y < 0 || n.y >= ly || n.x < 0 || n.x >= lx {
				continue
			}
			newCell := g[n.y][n.x]
			if int(newCell.height())-int(cell.height()) > 1 {
				continue
			}
			queue = append(queue, queueItem{
				pos:   n,
				steps: q.steps + 1,
			})
		}
	}
	return -1
}

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var grid grid

	var y int
	var queue1, queue2 []queueItem
	for sc.Scan() {
		row := make([]*cell, len(sc.Bytes()))
		for x, char := range sc.Bytes() {
			row[x] = &cell{char: char}
			if row[x].char == 'S' {
				queue1 = append(queue1, queueItem{pos: pos{y: y, x: x}, steps: 0})
			}
			if row[x].height() == 'a' {
				queue2 = append(queue2, queueItem{pos: pos{y: y, x: x}, steps: 0})
			}
		}

		grid = append(grid, row)
		y++
	}

	steps1 := shortestPath(grid, queue1)
	grid.reset()
	steps2 := shortestPath(grid, queue2)

	return []int{steps1, steps2}, nil
}
