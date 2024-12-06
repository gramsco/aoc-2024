package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	puzzle := newPuzzleFromInput()

	guardPath := getGuardPath(puzzle)
	fmt.Printf("Part1 : %d\n", len(getGuardPath(puzzle)))

	count := 0
	for _, coords := range guardPath {
		if isLoopyObstacle(puzzle, coords) {
			count++
		}
	}
	fmt.Printf("Part2: %d\n", count)

}

func newPuzzleFromInput() []string {

	file, err := os.Open("day6/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}

	return lines
}
