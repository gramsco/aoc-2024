package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct {
	LeftColumn  []int
	RightColumn []int
}

func (p *Puzzle) sortColumns() {
	sort.Slice(p.LeftColumn, func(i int, j int) bool {
		return p.LeftColumn[i] < p.LeftColumn[j]
	})
	sort.Slice(p.RightColumn, func(i int, j int) bool {
		return p.RightColumn[i] < p.RightColumn[j]
	})
}

func newPuzzleFromInput() Puzzle {

	file, err := os.Open("day1/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	puzzle := Puzzle{
		LeftColumn:  []int{},
		RightColumn: []int{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		splitted := strings.Split(txt, "   ")
		leftValue, err := strconv.Atoi(splitted[0])

		if err != nil {
			panic(err)
		}

		rightValue, err := strconv.Atoi(splitted[1])

		if err != nil {
			panic(err)
		}
		puzzle.LeftColumn = append(puzzle.LeftColumn, leftValue)
		puzzle.RightColumn = append(puzzle.RightColumn, rightValue)
	}

	return puzzle
}
