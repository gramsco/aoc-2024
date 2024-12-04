package main

import (
	"strings"
)

func solvePart1(s string) int {
	lines := strings.Split(s, "\n")
	count := 0

	for i, line := range lines {
		for j, char := range line {

			if char == 'X' {

				count += topLeft(lines, i, j) +
					top(lines, i, j) +
					topRight(lines, i, j) +
					right(lines, i, j) +
					bottomRight(lines, i, j) +
					bottom(lines, i, j) +
					bottomLeft(lines, i, j) +
					left(lines, i, j)

			}
		}
	}
	return count
}

const XMAS = "XMAS"

func topLeft(puzzle []string, topIndex int, leftIndex int) int {

	if topIndex < 3 {
		return 0
	}

	if leftIndex < 3 {
		return 0
	}

	x := string(puzzle[topIndex-0][leftIndex-0])
	m := string(puzzle[topIndex-1][leftIndex-1])
	a := string(puzzle[topIndex-2][leftIndex-2])
	s := string(puzzle[topIndex-3][leftIndex-3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func top(puzzle []string, topIndex int, leftIndex int) int {
	if topIndex < 3 {
		return 0
	}
	x := string(puzzle[topIndex-0][leftIndex])
	m := string(puzzle[topIndex-1][leftIndex])
	a := string(puzzle[topIndex-2][leftIndex])
	s := string(puzzle[topIndex-3][leftIndex])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func topRight(puzzle []string, topIndex int, leftIndex int) int {

	if topIndex < 3 {
		return 0
	}
	if leftIndex > len(puzzle[topIndex])-4 {
		return 0
	}
	x := string(puzzle[topIndex-0][leftIndex+0])
	m := string(puzzle[topIndex-1][leftIndex+1])
	a := string(puzzle[topIndex-2][leftIndex+2])
	s := string(puzzle[topIndex-3][leftIndex+3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func right(puzzle []string, topIndex int, leftIndex int) int {
	if leftIndex+3 >= len(puzzle[0]) {
		return 0
	}

	x := string(puzzle[topIndex][leftIndex])
	m := string(puzzle[topIndex][leftIndex+1])
	a := string(puzzle[topIndex][leftIndex+2])
	s := string(puzzle[topIndex][leftIndex+3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func bottomRight(puzzle []string, topIndex int, leftIndex int) int {

	if leftIndex+3 >= len(puzzle[0]) {
		return 0
	}

	if topIndex+3 >= len(puzzle) {
		return 0
	}

	x := string(puzzle[topIndex+0][leftIndex+0])
	m := string(puzzle[topIndex+1][leftIndex+1])
	a := string(puzzle[topIndex+2][leftIndex+2])
	s := string(puzzle[topIndex+3][leftIndex+3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func bottom(puzzle []string, topIndex int, leftIndex int) int {
	if topIndex+3 >= len(puzzle) {
		return 0
	}
	x := string(puzzle[topIndex+0][leftIndex+0])
	m := string(puzzle[topIndex+1][leftIndex+0])
	a := string(puzzle[topIndex+2][leftIndex+0])
	s := string(puzzle[topIndex+3][leftIndex+0])

	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func bottomLeft(puzzle []string, topIndex int, leftIndex int) int {

	if topIndex+3 >= len(puzzle) {
		return 0
	}

	if leftIndex < 3 {
		return 0
	}

	x := string(puzzle[topIndex+0][leftIndex-0])
	m := string(puzzle[topIndex+1][leftIndex-1])
	a := string(puzzle[topIndex+2][leftIndex-2])
	s := string(puzzle[topIndex+3][leftIndex-3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}

func left(puzzle []string, topIndex int, leftIndex int) int {
	if leftIndex < 3 {
		return 0
	}

	x := string(puzzle[topIndex+0][leftIndex-0])
	m := string(puzzle[topIndex+0][leftIndex-1])
	a := string(puzzle[topIndex+0][leftIndex-2])
	s := string(puzzle[topIndex+0][leftIndex-3])
	if x+m+a+s == XMAS {
		return 1
	}
	return 0
}
