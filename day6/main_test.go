package main

import (
	"strings"
	"testing"
)

func TestGuard(t *testing.T) {
	puzzle := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := getGuardPath(strings.Split(puzzle, "\n"))
	if len(result) != 41 {
		t.Fatalf("Expected 41, received %d", len(result))
	}

}
