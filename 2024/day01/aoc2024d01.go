package day01

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func readInputDat(inputDatPath string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	// Ensure the file can be opened
	inputDat, err := os.ReadFile(inputDatPath)
	if err != nil {
		return nil, nil, err
	}

	inputDatLines := strings.Split(string(inputDat), "\n")

	// Read the file line by line, and populate int arrays
	//inputDatScanner := bufio.NewScanner(inputDat)
	//for inputDatScanner.Scan() {

	for _, lineString := range inputDatLines {
		// parse the input.dat line and append slices
		var lefti int
		var righti int

		_, err := fmt.Sscanf(lineString, "%d   %d", &lefti, &righti)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}
		left = append(left, lefti)
		right = append(right, righti)
	}

	return left, right, nil
}

func calcDistances(left []int, right []int) []int {

	if len(left) != len(right) {
		return nil
	}

	sort.Ints(left)
	sort.Ints(right)

	distances := make([]int, len(left))
	for i := range len(left) {
		distances[i] = absDiffInt(left[i], right[i])
	}

	return distances
}

func sumIntSlice(intSlice []int) int {
	intSum := 0
	for _, number := range intSlice {
		intSum += number
	}
	return intSum
}

func Solution() {
	left, right, err := readInputDat("./day01/input.dat")
	if err != nil {
		panic(err)
	}

	distances := calcDistances(left, right)
	totalDistance := sumIntSlice(distances)
	fmt.Printf("Day 01 Answer: %d", totalDistance)
}
