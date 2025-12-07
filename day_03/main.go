package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MichaelBo1/aoc2025/aocutils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <main.go> <input_file>")
	}
	input := os.Args[1]

	lines := aocutils.MustScanToStringSlice(input)
	banks := aocutils.StringSliceTo2DMatrix(lines)

	// Part 1
	part1 := 0
	for _, bank := range banks {
		maxJoltage := 0
		// Naive O(n^2) [for each left, check each right]
		for left, digit := range bank {
			for right := left + 1; right < len(bank); right += 1 {
				joltage := digit*10 + bank[right]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		// fmt.Println("joltage for bank", maxJoltage, bank)
		part1 += maxJoltage
	}
	fmt.Println("Part_1:", part1)
}
