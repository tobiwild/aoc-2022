package puzzle

import (
	"bufio"
	"io"
)

const MAX_SIZE = 14

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	// alternative: scan runes
	scanner.Split(bufio.ScanBytes)

	var bytes []byte
	var result []int

	var pos int
	for scanner.Scan() {
		pos++
		bytes = append(bytes, scanner.Bytes()[0])
		blen := pos
		if blen > MAX_SIZE {
			bytes = bytes[1:]
			blen = MAX_SIZE
		}

		lastXUnique := func(count int) bool {
			if blen < count {
				return false
			}
			byteMap := make(map[byte]struct{})
			for i := blen - 1; i >= blen-count; i-- {
				b := bytes[i]
				if _, ok := byteMap[b]; ok {
					return false
				}
				byteMap[b] = struct{}{}
			}
			return true
		}

		if len(result) == 0 && lastXUnique(4) {
			result = append(result, pos)
		}
		if len(result) == 1 && lastXUnique(14) {
			result = append(result, pos)
			break
		}
	}

	return result, nil
}
