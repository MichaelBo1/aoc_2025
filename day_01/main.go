package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MichaelBo1/aoc2025/aocutils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <main.go> <input_file>")
	}
	input := os.Args[1]

	rotations := aocutils.MustScanToStringSlice(input)
	p1Answer, p2Answer := answer(rotations)
	fmt.Println("PART_1:", p1Answer)
	fmt.Println("PART_2:", p2Answer)
}

func answer(rotations []string) (part1, part2 int) {
	part1, part2 = 0, 0
	// Starts at 50 per https://adventofcode.com/2025/day/1
	dial := 50
	for _, rot := range rotations {
		direction, amountStr := rot[0], rot[1:]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(err)
		}

		part2 += countZeroCrosses(dial, direction, amount)

		if direction == 'L' {
			amount *= -1
		}

		// 0 - 99 and circuar.
		dial = mod(dial+amount, 100)
		if dial == 0 {
			part1 += 1
		}
	}

	return
}

func mod(a, n int) int {
	return (a%n + n) % n
}

func countZeroCrosses(start int, dir byte, steps int) int {
	var firstCross int
	if dir == 'R' {
		firstCross = (100 - start) % 100
	} else {
		firstCross = start % 100
	}

	if firstCross == 0 {
		firstCross = 100
	}

	if steps < firstCross {
		return 0
	}

	return 1 + (steps-firstCross)/100

}
