package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MichaelBo1/aoc2025/aocutils"
)

func stringSliceToDigits(list []string) [][]int {

}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <main.go> <input_file>")
	}
	input := os.Args[1]

	banks := aocutils.MustScanToStringSlice(input)
	fmt.Println(banks)
}
