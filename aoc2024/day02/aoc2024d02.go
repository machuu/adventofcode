package day02

import (
	"encoding/json"
	"fmt"

	"github.com/machuu/adventofcode/aoc2024/utils"
)

func countSafeReports(reports [][]int) int {
	safeReports := 0
	for j, report := range reports {
		safe := true
		increased := false
		decreased := false

		// get report as string
		reportJson, _ := json.Marshal(report)
		reportString := string(reportJson)
		utils.Use(j, reportString)
		for i := range len(report) {
			if i == len(report)-1 {
				break
			}

			currentLevel := report[i]
			nextLevel := report[i+1]
			//fmt.Printf("Checking levels [%d]: %d -> %d\n", i, report[i], report[i+1])
			diff := utils.AbsDiffInt(currentLevel, nextLevel)

			if currentLevel > nextLevel {
				decreased = true
			} else {
				increased = true
			}

			if increased && decreased {
				//fmt.Printf("Report [%d] is unsafe: inscreased and decreased in %s\n", j, reportString)
				safe = false
				break
			} else if diff > 3 {
				//fmt.Printf("Report [%d] is unsafe: |%d-%d| > 3\n", j, report[i], report[i+1])
				safe = false
				break
			} else if diff == 0 {
				//fmt.Printf("Report [%d] is unsafe: |%d-%d| = 0\n", j, report[i], report[i+1])
				safe = false
				break
			}
		}
		if safe {
			safeReports += 1
		}
	}
	return safeReports
}

func SolnPart1() {
	reports, err := utils.ParseInputInto2DIntSlice("./day02/input.dat")
	if err != nil {
		panic(err)
	}

	safeReports := countSafeReports(reports)
	fmt.Printf("Day 02 Part 1 Answer: %d", safeReports)
}

func Solution() {
	SolnPart1()
}
