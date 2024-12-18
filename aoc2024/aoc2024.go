package main

import (
	"os"
	"slices"

	"github.com/machuu/adventofcode/aoc2024/day01"
	"github.com/machuu/adventofcode/aoc2024/day02"
	"github.com/machuu/adventofcode/aoc2024/day03"
)

func main() {
	args := os.Args[1:]

	if slices.Contains(args, "day01") {
		day01.Solution()
	} else if slices.Contains(args, "day02") {
		day02.Solution()
	} else if slices.Contains(args, "day03") {
		day03.Solution()
	} else {
		day01.Solution()
		day02.Solution()
		day03.Solution()
	}
}
