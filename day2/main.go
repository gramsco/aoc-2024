package main

import (
	"fmt"
	"math"
)

func main() {
	puzzle := getListOfReports()
	result := SolvePart1(puzzle)
	fmt.Printf("First result is %d\n", result)

	result2 := SolvePart2(puzzle)
	fmt.Printf("Second result is %d\n", result2)
}

func SolvePart1(puzzle [][]int) int {
	count := 0
	for _, report := range puzzle {
		if CheckReportSafety(report) {
			count += 1
		}
	}
	return count
}

func SolvePart2(puzzle [][]int) int {
	count := 0
	for _, report := range puzzle {
		if CheckReportSafetyWithDampener(report) {
			count += 1
		}
	}
	return count
}

func CheckReportSafety(report []int) bool {
	sign := math.Signbit(float64(report[len(report)-1] - report[0]))

	for currentLevelIndex := 0; currentLevelIndex < len(report)-1; currentLevelIndex++ {
		nextLevelIndex := currentLevelIndex + 1

		distance := float64(report[nextLevelIndex] - report[currentLevelIndex])

		if math.Signbit(distance) != sign {
			return false
		}

		switch math.Abs(distance) {
		case 1:
		case 2:
		case 3:
			continue
		default:
			return false
		}
	}
	return true
}

func CheckReportSafetyWithDampener(reports []int) bool {

	if CheckReportSafety(reports) {
		return true
	}

	for index := range reports {
		if CheckReportSafety(removeLevelByindex(reports, index)) {
			return true
		}
	}

	return false

}
