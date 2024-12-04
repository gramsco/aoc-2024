package main

import (
	"fmt"
)

func main() {
	txt := newPuzzleFromInput()
	fmt.Printf("Part1: %d\n", solvePart1(txt))
	fmt.Printf("Part2: %d\n", solvePart2(txt))
}
