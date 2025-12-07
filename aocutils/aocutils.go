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
