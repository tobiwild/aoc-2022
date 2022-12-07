package puzzle

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const DISK_SPACE = 70000000
const REQUIRED_FREE_DISK_SPACE = 30000000

func getDirSizes(r io.Reader) map[string]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	dirSizes := make(map[string]int)

	var dirStack []string
	var curSize int
	var curDir string

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fmt.Sscanf(line, "$ cd %v", &curDir); err == nil {
			if curDir == "/" {
				dirStack = nil
			} else if curDir == ".." {
				if len(dirStack) > 0 {
					dirStack = dirStack[:len(dirStack)-1]
				}
			} else {
				dirStack = append(dirStack, curDir)
			}
			continue
		}

		if _, err := fmt.Sscanf(line, "%d", &curSize); err == nil {
			for i := 0; i <= len(dirStack); i++ {
				key := "/" + strings.Join(dirStack[:i], "/")
				dirSizes[key] += curSize
			}
		}
	}

	return dirSizes
}

func Solve(r io.Reader) ([]int, error) {
	dirSizes := getDirSizes(r)

	var result1 int
	for _, size := range dirSizes {
		if size <= 100000 {
			result1 += size
		}
	}

	usedSpace := dirSizes["/"]
	unusedSpace := DISK_SPACE - usedSpace
	minDeleteSize := REQUIRED_FREE_DISK_SPACE - unusedSpace

	var deleteSize int
	for _, size := range dirSizes {
		if size < minDeleteSize {
			continue
		}
		if deleteSize == 0 || size < deleteSize {
			deleteSize = size
		}
	}

	return []int{result1, deleteSize}, nil
}
