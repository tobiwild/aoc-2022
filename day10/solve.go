package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type crt struct {
	i               int
	width           int
	pixels          []bool // true = lit
	middleSpritePos int
	spriteSize      int
}

type result struct {
	signalStrength int
	screenImage    string
}

func (c *crt) setPixel() {
	min := c.middleSpritePos - c.spriteSize/2
	max := c.middleSpritePos + c.spriteSize/2
	i := c.i
	rowI := i % c.width
	if rowI >= min && rowI <= max {
		c.pixels[i] = true
	}
	c.i++
}

func (c crt) draw(w io.Writer) {
	for i, pixel := range c.pixels {
		if i > 0 && i%c.width == 0 {
			fmt.Fprintln(w)
		}
		if pixel {
			fmt.Fprint(w, "#")
		} else {
			fmt.Fprint(w, ".")
		}
	}
	fmt.Fprintln(w)
}

func Solve(r io.Reader) (*result, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	x := 1
	var v, cycle, signalStrength int

	crt := crt{
		width:           40,
		pixels:          make([]bool, 240),
		middleSpritePos: 1,
		spriteSize:      3,
	}

	wait := func(c int) {
		for i := 0; i < c; i++ {
			cycle++
			crt.setPixel()
			if cycle%40 == 20 {
				signalStrength += cycle * x
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "noop" {
			wait(1)
			continue
		}

		if _, err := fmt.Sscanf(line, "addx %d", &v); err != nil {
			return nil, err
		}

		wait(2)
		x += v
		crt.middleSpritePos = x
	}

	// crt.draw(os.Stdout)

	var screenImage strings.Builder
	fmt.Fprintln(&screenImage)
	crt.draw(&screenImage)

	return &result{
		signalStrength: signalStrength,
		screenImage:    screenImage.String(),
	}, nil
}
