package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	puzzle := readPuzzle()
	fmt.Println(part1(puzzle))
	fmt.Println(part2(puzzle))
}

func part1(s string) int {
	puzzle := strings.Split(s, "\n")
	sum := 0
	for y := range puzzle {
		for x := range len(puzzle[y]) {
			head := string(puzzle[y][x])
			if head == "0" {
				found := []Coord{}
				visit(puzzle, Coord{y, x}, 1, &found, false)
				sum += len(found)
			}
		}
	}
	return sum
}

func part2(s string) int {
	puzzle := strings.Split(s, "\n")
	sum := 0
	for y := range puzzle {
		for x := range len(puzzle[y]) {
			head := string(puzzle[y][x])
			if head == "0" {
				found := []Coord{}
				visit(puzzle, Coord{y, x}, 1, &found, true)
				sum += len(found)
			}
		}
	}
	return sum
}

func visit(puzzle []string, currentCoordinates Coord, target int, peaksFound *[]Coord, isPart2 bool) *[]Coord {

	if currentCoordinates.isNil() {
		return nil
	}

	if target > 9 {
		f := *peaksFound
		if isPart2 {
			*peaksFound = append(*peaksFound, currentCoordinates)
			return nil
		}
		if !slices.Contains(f, currentCoordinates) {
			*peaksFound = append(*peaksFound, currentCoordinates)
		}
		return nil
	}

	visit(puzzle, Down(puzzle, currentCoordinates, target), target+1, peaksFound, isPart2)
	visit(puzzle, Left(puzzle, currentCoordinates, target), target+1, peaksFound, isPart2)
	visit(puzzle, Top(puzzle, currentCoordinates, target), target+1, peaksFound, isPart2)
	visit(puzzle, Right(puzzle, currentCoordinates, target), target+1, peaksFound, isPart2)

	return peaksFound

}

func readPuzzle() string {

	file, err := os.Open("day10/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n")

}
