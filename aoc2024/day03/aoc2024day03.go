package day03

import (
	"fmt"

	"github.com/machuu/adventofcode/aoc2024/utils"
)

type mul struct {
	num1 int
	num2 int
}

func mulStringsToMuls(mulStrings []string) []mul {
	var muls []mul
	for _, mulString := range mulStrings {
		var num1 int
		var num2 int
		fmt.Sscanf(mulString, "mul(%d,%d)", &num1, &num2)
		fmt.Printf("extracting mul: '%s' -> (%d,%d)\n", mulString, num1, num2)
		muls = append(muls, mul{num1, num2})
	}

	return muls
}

func multiplyMuls(muls []mul) int {
	result := 0

	for _, mul := range muls {
		result += mul.num1 * mul.num2
	}

	return result
}

func SolnPart1() {
	mulStrings := utils.GetRegexMatches(`mul\(\d{1,3},\d{1,3}\)`, "./day03/input.dat")
	muls := mulStringsToMuls(mulStrings)
	results := multiplyMuls(muls)
	fmt.Printf("Day 03 Part 1: %d", results)
}

func Solution() {
	SolnPart1()
}
