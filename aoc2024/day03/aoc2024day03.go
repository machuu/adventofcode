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
	do := true
	var command string
	for _, mulString := range mulStrings {
		var num1 int
		var num2 int
		nNums, _ := fmt.Sscanf(mulString, "mul(%d,%d)", &num1, &num2)
		if nNums == 0 {
			nCommand, _ := fmt.Sscanf(mulString, "%s", &command)
			if nCommand == 0 {
				// shrug
			} else if command == "do()" {
				do = true
			} else if command == "don't()" {
				do = false
			} else {
				fmt.Printf("Unrecognized command: %s\n", command)
			}
			continue
		}
		//fmt.Printf("extracting mul: '%s' -> (%d,%d)\n", mulString, num1, num2)
		if do {
			muls = append(muls, mul{num1, num2})
		}
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
	fmt.Printf("Day 03 Part 1: %d\n", results)
}

func SolnPart2() {
	mulStrings := utils.GetRegexMatches(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`, "./day03/input.dat")
	muls := mulStringsToMuls(mulStrings)
	results := multiplyMuls(muls)
	fmt.Printf("Day 03 Part 2: %d\n", results)
}

func Solution() {
	SolnPart1()
	SolnPart2()
}
