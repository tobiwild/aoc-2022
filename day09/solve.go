package puzzle

import (
	"bufio"
	"fmt"
	"io"
)

type pos struct {
	x, y int
}

func (p1 *pos) add(p2 pos) {
	p1.x += p2.x
	p1.y += p2.y
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func sign(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}

type knot struct {
	pos         pos
	visited     map[pos]struct{}
	recordMoves bool
}

func newKnot() knot {
	p := pos{}
	visited := make(map[pos]struct{})
	visited[p] = struct{}{}
	return knot{
		pos:     p,
		visited: visited,
	}
}

func (k *knot) move(p pos) {
	k.pos.add(p)
	if k.recordMoves {
		k.visited[k.pos] = struct{}{}
	}
}

func (k1 *knot) adjustTo(k2 knot) {
	dx := k2.pos.x - k1.pos.x
	dy := k2.pos.y - k1.pos.y
	adx := abs(dx)
	ady := abs(dy)
	if adx > 1 && ady > 1 {
		k1.move(pos{
			x: (adx - 1) * sign(dx),
			y: (ady - 1) * sign(dy),
		})
	} else if adx > 1 {
		k1.move(pos{
			x: (adx - 1) * sign(dx),
			y: dy,
		})
	} else if ady > 1 {
		k1.move(pos{
			x: dx,
			y: (ady - 1) * sign(dy),
		})
	}
}

type rope []knot

func newRope(knotCount int) rope {
	rope := make([]knot, 0, knotCount)
	for i := 0; i < knotCount; i++ {
		rope = append(rope, newKnot())
	}
	return rope
}

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var dir byte
	var steps int

	rope := newRope(10)
	rope[1].recordMoves = true
	rope[9].recordMoves = true

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fmt.Sscanf(line, "%c %d", &dir, &steps); err != nil {
			return nil, err
		}

		var headDiff pos
		switch dir {
		case 'R':
			headDiff.x = 1
		case 'L':
			headDiff.x = -1
		case 'U':
			headDiff.y = 1
		case 'D':
			headDiff.y = -1
		}

		for i := 0; i < steps; i++ {
			for r := 0; r < len(rope)-1; r++ {
				head := &rope[r]
				tail := &rope[r+1]
				if r == 0 {
					head.move(headDiff)
				}
				tail.adjustTo(*head)
			}
		}
	}

	result1 := len(rope[1].visited)
	result2 := len(rope[9].visited)

	return []int{result1, result2}, nil
}
