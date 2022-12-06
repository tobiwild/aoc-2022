package puzzle

import (
	"bufio"
	"io"
)

func Solve(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)

	var bytes []byte
	var result1, result2 *int

	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes()[0])
		blen := len(bytes)

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

		if result1 == nil && lastXUnique(4) {
			result1 = &blen
		}
		if result2 == nil && lastXUnique(14) {
			result2 = &blen
		}
		if result1 != nil && result2 != nil {
			return []int{*result1, *result2}, nil
		}
	}

	return []int{}, nil
}
