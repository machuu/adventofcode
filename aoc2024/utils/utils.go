package utils

import (
	"fmt"
	"io"
	"os"
	"regexp"
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

func ParseInputInto2DIntSlice(inputDatPath string) ([][]int, error) {
	inputDatLines, _ := ReadInputDat(inputDatPath)

	lineSlices := [][]int{}
	// Turn each line into a slice of ints
	for _, lineString := range inputDatLines {
		//fmt.Printf("Parsing Line: '%s'\n", lineString)
		if lineString == "" {
			break
		}
		lineSlice := []int{}
		for _, intStr := range strings.Split(lineString, " ") {
			var parsedInt int
			_, err := fmt.Sscanf(intStr, "%d", &parsedInt)
			if err == io.EOF {
				break
			}
			//fmt.Printf("Appending: '%d'\n", parsedInt)
			lineSlice = append(lineSlice, parsedInt)

		}
		//fmt.Printf("lineSlice: %s\n", lineSlice)
		lineSlices = append(lineSlices, lineSlice)
	}

	//fmt.Printf("lineSlices: %s\n", lineSlices)
	return lineSlices, nil
}

func GetRegexMatches(regexPatternString string, inputDatPath string) []string {
	inputDatLines, _ := ReadInputDat(inputDatPath)
	regexPattern := regexp.MustCompile(regexPatternString)

	var regexMatches []string
	for _, lineString := range inputDatLines {
		//
		regexMatches = append(regexMatches, regexPattern.FindAllString(lineString, -1)...)
	}

	return regexMatches
}

func SumIntSlice(intSlice []int) int {
	intSum := 0
	for _, number := range intSlice {
		intSum += number
	}
	return intSum
}

func PopIndexFromIntSlice(index int, sourceIntSlice []int) []int {
	// append with nil slice to create a copy of sourceIntList
	// appending original slice overwrite original slice in parent func
	intSlice := append([]int(nil), sourceIntSlice...)
	poppedIntSlice := append(intSlice[:index], intSlice[index+1:]...)
	return poppedIntSlice
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

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
