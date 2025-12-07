package aocutils

import (
	"bufio"
	"os"
	"strings"
)

// MustScanToStringSlice scans a file line by line into a []string.
// It panics if it fails to open the file.
func MustScanToStringSlice(filename string) []string {
	fhand, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var result []string

	scan := bufio.NewScanner(fhand)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		result = append(result, line)
	}

	return result
}

// StringSliceTo2DMatrix converts []string of large integers into [][]int,
// where each row is a slice of the individual digits.
func StringSliceTo2DMatrix(list []string) [][]int {
	result := make([][]int, len(list))

	for i, s := range list {
		row := make([]int, len(s))
		for j, ch := range s {
			row[j] = int(ch - '0')
		}
		result[i] = row
	}

	return result
}
