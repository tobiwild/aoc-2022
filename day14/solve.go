package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type pos struct {
	x, y int
}

type grid map[int]map[int]struct{}

type cave struct {
	g       grid
	bottomY int
}

func (c cave) add(p pos) {
	if _, ok := c.g[p.x]; !ok {
		c.g[p.x] = make(map[int]struct{})
	}
	c.g[p.x][p.y] = struct{}{}
}

func (c cave) getMaxY() int {
	var max int
	for _, col := range c.g {
		for y := range col {
			if y > max {
				max = y
			}
		}
	}
	return max
}

func (c cave) addLine(p1, p2 pos) {
	if p1.x == p2.x {
		min := min(p1.y, p2.y)
		max := max(p1.y, p2.y)
		for y := min; y <= max; y++ {
			c.add(pos{x: p1.x, y: y})
		}
	} else {
		min := min(p1.x, p2.x)
		max := max(p1.x, p2.x)
		for x := min; x <= max; x++ {
			c.add(pos{x: x, y: p1.y})
		}
	}
}

func (c cave) blockedAt(p pos) bool {
	if c.bottomY > 0 && p.y >= c.bottomY {
		return true
	}
	_, ok := c.g[p.x][p.y]
	return ok
}

func (c cave) insertable(p pos) bool {
	if c.bottomY > 0 {
		if _, ok := c.g[500][0]; ok {
			return false
		}
		return true
	}
	col, ok := c.g[p.x]
	if !ok {
		return false
	}

	for y := range col {
		if y > p.y {
			return true
		}
	}

	return false
}

func (c cave) insertSand() bool {
	current := pos{x: 500, y: 0}
	for c.insertable(current) {
		next := current

		next.y++
		if !c.blockedAt(next) {
			current = next
			continue
		}

		next.x--
		if !c.blockedAt(next) {
			current = next
			continue
		}

		next.x += 2
		if !c.blockedAt(next) {
			current = next
			continue
		}

		c.add(current)
		return true
	}

	return false
}

func Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	cave1 := cave{
		g: make(grid),
	}
	cave2 := cave{
		g: make(grid),
	}
	for sc.Scan() {
		var lastPos *pos
		for _, posStr := range strings.Split(sc.Text(), " -> ") {
			var pos pos
			if _, err := fmt.Sscanf(posStr, "%d,%d", &pos.x, &pos.y); err != nil {
				return nil, err
			}
			if lastPos != nil {
				cave1.addLine(*lastPos, pos)
				cave2.addLine(*lastPos, pos)
			}
			lastPos = &pos
		}
	}

	var unitCount1 int
	for cave1.insertSand() {
		unitCount1++
	}

	cave2.bottomY = cave2.getMaxY() + 2
	var unitCount2 int
	for cave2.insertSand() {
		unitCount2++
	}

	return []int{unitCount1, unitCount2}, nil
}
