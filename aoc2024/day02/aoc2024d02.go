package day02

import (
	"encoding/json"
	"fmt"

	"github.com/machuu/adventofcode/aoc2024/utils"
)

func countSafeReports(reports [][]int) int {
	safeReports := 0
	for j, report := range reports {
		reportJson, _ := json.Marshal(report)
		reportString := string(reportJson)
		utils.Use(j, reportString)
		//fmt.Printf("\nChecking if report is safe: %s\n", reportString)

		if isReportSafe(report, false) {
			safeReports += 1
		}
	}
	return safeReports
}

func countSafeReportsWithProblemDampener(reports [][]int) int {
	safeReports := 0
	for j, report := range reports {
		reportJson, _ := json.Marshal(report)
		reportString := string(reportJson)
		utils.Use(j, reportString)
		//fmt.Printf("\nChecking if report is safe: %s\n", reportString)

		if isReportSafe(report, true) {
			safeReports += 1
		}
	}
	return safeReports
}

func isReportSafe(report []int, adjustForPD bool) bool {
	// get report as string
	reportJson, _ := json.Marshal(report)
	reportString := string(reportJson)
	utils.Use(reportString)

	safe := true
	increasedEver := false
	decreasedEver := false

	for i, level := range report {
		utils.Use(level)
		var currentLevel int
		var nextLevel int
		var diff int
		lastLevel := false
		decreased := false
		increased := false

		currentLevel = report[i]
		if i == len(report)-1 {
			lastLevel = true
		} else {
			nextLevel = report[i+1]
			diff = utils.AbsDiffInt(currentLevel, nextLevel)
		}

		//fmt.Printf("Checking levels [%d]: %d -> %d\n", i, currentLevel, nextLevel)

		if lastLevel {
			decreased = false
			increased = false
		} else if currentLevel > nextLevel {
			decreasedEver = true
			decreased = true
		} else {
			increasedEver = true
			increased = true
		}

		if lastLevel {
			// don't do anything, this is just so adjustForPD can
			// check if removing the last level fixes the report
		} else if decreased && increasedEver {
			//fmt.Printf("Report is unsafe: decrease after increase %d->%d -> %s\n", currentLevel, nextLevel, reportString)
			safe = false
		} else if increased && decreasedEver {
			//fmt.Printf("Report is unsafe: increase after decrease %d->%d -> %s\n", currentLevel, nextLevel, reportString)
			safe = false
		} else if diff > 3 {
			//fmt.Printf("Report is unsafe: |%d-%d| > 3 -> %s\n", currentLevel, nextLevel, reportString)
			safe = false
		} else if diff == 0 {
			//fmt.Printf("Report is unsafe: |%d-%d| = 0 -> %s\n", currentLevel, nextLevel, reportString)
			safe = false
		}

		// have to check for every level on first pass, since
		// removing the first level could be the one that makes the report safe
		if adjustForPD {
			// check if modified reports are safe
			// report without currentLevel
			adjustedReport := utils.PopIndexFromIntSlice(i, report)
			adjustedReportJson, _ := json.Marshal(adjustedReport)
			adjustedReportString := string(adjustedReportJson)
			utils.Use(adjustedReportString)
			//fmt.Printf("Checking adjusted report removing index %d (%d): %s -> %s\n", i, level, reportString, adjustedReportString)
			safe = isReportSafe(adjustedReport, false)
			if safe {
				//fmt.Printf("Success\n")
				//fmt.Printf("Fixed unsafe report by removing index %d (%d)\n%s\n%s\n\n", i, report[i], reportString, adjustedReportString)
				break
			} else if lastLevel {
				//fmt.Printf("unfixable report: %s\n", reportString)
			} else {
				//fmt.Printf("Failed\n")
				//fmt.Printf("Unable to fix by removing index %d (%d): %s -> %s\n", i, level, reportString, adjustedReportString)
				continue
			}
		} else if !safe {
			break
		}
	}

	return safe
}

func SolnPart1(reports [][]int) {
	safeReports := countSafeReports(reports)
	fmt.Printf("Day 02 Part 1 Answer: %d\n", safeReports)
}

func SolnPart2(reports [][]int) {
	safeReports := countSafeReportsWithProblemDampener(reports)
	fmt.Printf("Day 02 Part 2 Answer: %d\n", safeReports)
}

func Solution() {
	reports, err := utils.ParseInputInto2DIntSlice("./day02/input.dat")
	if err != nil {
		panic(err)
	}

	SolnPart1(reports)
	SolnPart2(reports)
}
