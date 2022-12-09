package puzzle

import (
	"bufio"
	"fmt"
	"io"
)

type pos struct {
	x, y int
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func sig(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}

type rope struct {
	knots   []pos
	tailMap map[pos]struct{}
}

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var dir string
	var steps int

	rope1 := rope{
		knots:   make([]pos, 2),
		tailMap: make(map[pos]struct{}),
	}
	rope2 := rope{
		knots:   make([]pos, 10),
		tailMap: make(map[pos]struct{}),
	}

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fmt.Sscanf(line, "%v %v", &dir, &steps); err != nil {
			return nil, err
		}

		var headDX, headDY int
		switch dir {
		case "R":
			headDX = 1
		case "L":
			headDX = -1
		case "U":
			headDY = 1
		case "D":
			headDY = -1
		}

		updateRope := func(rope *rope) {
			for i := 0; i < steps; i++ {
				for r := 0; r < len(rope.knots)-1; r++ {
					head := &rope.knots[r]
					tail := &rope.knots[r+1]
					if r == 0 {
						*head = pos{
							x: head.x + headDX,
							y: head.y + headDY,
						}
					}
					dx := head.x - tail.x
					dy := head.y - tail.y
					adx := abs(dx)
					ady := abs(dy)
					var tailDX, tailDY int
					if adx > 1 && ady > 1 {
						tailDX = (adx - 1) * sig(dx)
						tailDY = (ady - 1) * sig(dy)
					} else if adx > 1 {
						tailDX = (adx - 1) * sig(dx)
						if ady > 0 {
							tailDY = dy
						}
					} else if ady > 1 {
						tailDY = (ady - 1) * sig(dy)
						if adx > 0 {
							tailDX = dx
						}
					}
					*tail = pos{
						x: tail.x + tailDX,
						y: tail.y + tailDY,
					}
					if r == len(rope.knots)-2 {
						rope.tailMap[*tail] = struct{}{}
					}
				}
			}
		}

		updateRope(&rope1)
		updateRope(&rope2)
	}

	result1 := len(rope1.tailMap)
	result2 := len(rope2.tailMap)

	return []int{result1, result2}, nil
}
