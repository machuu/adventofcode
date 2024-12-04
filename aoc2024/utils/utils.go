package utils

import (
	"os"
	"strings"
)

func ReadInputDat(inputDatPath string) ([]string, error) {
	// Ensure the file can be opened
	inputDat, err := os.ReadFile(inputDatPath)
	if err != nil {
		return nil, err
	}

	inputDatLines := strings.Split(string(inputDat), "\n")
	return inputDatLines, nil
}

func SumIntSlice(intSlice []int) int {
	intSum := 0
	for _, number := range intSlice {
		intSum += number
	}
	return intSum
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
