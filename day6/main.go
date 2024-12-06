package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	puzzle := newPuzzleFromInput()

	guardPath := getGuardPath(puzzle)
	fmt.Printf("Part1 : %d\n", len(getGuardPath(puzzle)))

	var ops atomic.Uint32
	var wg sync.WaitGroup
	for _, coords := range guardPath {
		wg.Add(1)

		go func() {
			if isLoopyObstacle(puzzle, coords) {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Part2: %d\n", ops.Load())

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
