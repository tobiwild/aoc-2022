package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

type pos struct {
	x, y int
}

type pair struct {
	sensor, beacon pos
}

type pairs []pair

type rng struct {
	min, max int
}

type rngs []rng

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (r rngs) merge() rngs {
	sort.Slice(r, func(i, j int) bool {
		return r[i].min < r[j].min
	})
	var result rngs
	for i, rng := range r {
		if i == 0 {
			result = append(result, rng)
			continue
		}
		if rng.min > result[len(result)-1].max {
			result = append(result, rng)
		} else {
			result[len(result)-1].max = max(result[len(result)-1].max, rng.max)
		}
	}

	return result
}

func (r rng) length() int {
	result := r.max - r.min + 1
	if result < 0 {
		return 0
	}
	return result
}

func (rngs rngs) length() int {
	var result int
	for _, rng := range rngs {
		result += rng.length()
	}
	return result
}

func getRange(start, end pos, y int) rng {
	dx := abs(start.x - end.x)
	dy := abs(start.y - end.y)
	distance := dx + dy

	minX := start.x - distance + abs(start.y-y)
	maxX := start.x + distance - abs(start.y-y)

	return rng{min: minX, max: maxX}
}

func (p pair) beaconFreeRange(y int) rng {
	return getRange(p.sensor, p.beacon, y)
}

func (p pairs) beaconFreeRanges(y int) rngs {
	var result rngs
	for _, pair := range p {
		rng := pair.beaconFreeRange(y)
		if rng.length() == 0 {
			continue
		}
		result = append(result, rng)
	}
	return result.merge()
}

func (p pairs) beaconCount(y int) int {
	result := make(map[int]struct{})
	for _, pair := range p {
		if pair.beacon.y == y {
			result[pair.beacon.x] = struct{}{}
		}
	}
	return len(result)
}

type puzzle struct {
	part1Y   int
	part2Max int
}

func (p puzzle) Solve(r io.Reader) ([]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var pairs pairs

	for sc.Scan() {
		var pair pair
		if _, err := fmt.Sscanf(
			sc.Text(),
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pair.sensor.x,
			&pair.sensor.y,
			&pair.beacon.x,
			&pair.beacon.y,
		); err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}

	result1 := pairs.beaconFreeRanges(p.part1Y).length() - pairs.beaconCount(p.part1Y)

	min := 0
	max := p.part2Max
	getTuningFrequency := func() int {
		for y := min; y <= max; y++ {
			rngs := pairs.beaconFreeRanges(y)
			if len(rngs) == 0 {
				continue
			}
			if len(rngs) == 1 && rngs[0].min <= min && rngs[0].max >= max {
				continue
			}
			for x := min; x <= max; x++ {
				included := func() bool {
					for _, rng := range rngs {
						if x >= rng.min && x <= rng.max {
							return true
						}
					}
					return false
				}
				if !included() {
					return x*4000000 + y
				}
			}
		}
		return 0
	}

	result2 := getTuningFrequency()

	return []int{result1, result2}, nil
}
