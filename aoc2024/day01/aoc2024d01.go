package day01

import (
	"fmt"
	"io"
	"sort"

	"github.com/machuu/adventofcode/aoc2024/utils"
)

func readInputDat(inputDatPath string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	inputDatLines, err := utils.ReadInputDat(inputDatPath)
	if err != nil {
		return nil, nil, err
	}

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

	distances := make([]int, len(left))
	for i := range len(left) {
		distances[i] = utils.AbsDiffInt(left[i], right[i])
	}

	return distances
}

func getTotalDistanceBetweenLists(left []int, right []int) int {
	distances := calcDistances(left, right)
	totalDistance := utils.SumIntSlice(distances)
	return totalDistance
}

func countMatchesFromLeftInRight(left []int, right []int) map[int]int {
	// convert left and right into maps, so each
	countedMatches := make(map[int]int)
	for _, leftNum := range left {
		matches := 0
		for _, rightNum := range right {
			// lists are sorted, so break if rightNum > leftNum
			if rightNum > leftNum {
				break
			} else if leftNum == rightNum {
				matches += 1
			}
		}
		if matches > 0 {
			//fmt.Printf("{'%d': %d},\n", leftNum, matches)
			countedMatches[leftNum] = matches
		}
	}
	return countedMatches
}

func getSimilarityBetweenLists(left []int, right []int) int {
	countedMatchesFromLeftInRight := countMatchesFromLeftInRight(left, right)

	similarity := 0

	for id, count := range countedMatchesFromLeftInRight {
		similarity += id * count
	}

	return similarity
}

func Solution() {
	left, right, err := readInputDat("./day01/input.dat")
	if err != nil {
		panic(err)
	}

	// sort lists
	sort.Ints(left)
	sort.Ints(right)

	// Part 1
	totalDistance := getTotalDistanceBetweenLists(left, right)
	fmt.Printf("Day 01 Part 1 Answer: %d\n", totalDistance)

	// Part 2
	similarity := getSimilarityBetweenLists(left, right)
	fmt.Printf("Day 01 Part 2 Answer: %d\n", similarity)
}
