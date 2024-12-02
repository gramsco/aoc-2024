package main

import (
	"fmt"
	"math"
)

func main() {
	puzzle := newPuzzleFromInput()
	result1 := SolvePart1(&puzzle)
	fmt.Printf("Result1 = %d\n", result1)

	result2 := SolvePart2(&puzzle)
	fmt.Printf("Result2 = %d\n", result2)
}

func SolvePart1(puzzle *Puzzle) int {
	var sumDistances int
	puzzle.sortColumns()

	for i := range puzzle.LeftColumn {
		left_value := puzzle.LeftColumn[i]
		right_value := puzzle.RightColumn[i]
		to_add := math.Abs(float64(left_value) - float64(right_value))
		sumDistances += int(to_add)
	}
	return sumDistances
}

func SolvePart2(puzzle *Puzzle) int {
	count := count(puzzle.RightColumn)
	var result int
	for _, val := range puzzle.LeftColumn {
		result += val * count[val]
	}
	return result
}

func count(column []int) map[int]int {
	m := make(map[int]int)
	for _, v := range column {
		m[v] = m[v] + 1
	}
	return m
}
