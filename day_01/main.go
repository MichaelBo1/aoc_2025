package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mustScanToSlice(input string) []string {
	fhand, err := os.Open(input)
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

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <main.go> <input_file>")
	}
	input := os.Args[1]

	rotations := mustScanToSlice(input)
	p1Answer := part1(rotations)
	fmt.Println("PART_1:", p1Answer)
}

func part1(rotations []string) int {
	result := 0
	// Starts at 50 per https://adventofcode.com/2025/day/1
	dial := 50
	for _, rot := range rotations {
		direction, amountStr := rot[0], rot[1:]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			amount *= -1
		}
		// 0 - 99 and circuar.
		dial = mod(dial+amount, 100)
		if dial == 0 {
			result += 1
		}
	}

	return result
}

func mod(a, n int) int {
	return (a%n + n) % n
}
