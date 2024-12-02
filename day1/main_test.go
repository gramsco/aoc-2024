package main

import (
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	puzzle1 := Puzzle{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}

	result := SolvePart1(&puzzle1)
	if result != 11 {
		t.Fatalf("Expected %d, received %d\n", 11, result)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	puzzle1 := Puzzle{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}

	result := SolvePart2(&puzzle1)
	if result != 31 {
		t.Fatalf("Expected %d, received %d\n", 31, result)
	}
}

func TestSortPuzzle(t *testing.T) {
	puzzle1 := Puzzle{
		LeftColumn:  []int{3, 4, 2, 1, 3, 3},
		RightColumn: []int{4, 3, 5, 3, 9, 3},
	}

	expected := Puzzle{
		LeftColumn:  []int{1, 2, 3, 3, 3, 4},
		RightColumn: []int{3, 3, 3, 4, 5, 9},
	}
	puzzle1.sortColumns()

	for i := range puzzle1.LeftColumn {
		if puzzle1.LeftColumn[i] != expected.LeftColumn[i] {
			t.Fatalf("Expected %d, received %d", puzzle1.LeftColumn[i], expected.LeftColumn[i])
		}
		if puzzle1.RightColumn[i] != expected.RightColumn[i] {
			t.Fatalf("Expected %d, received %d", puzzle1.RightColumn[i], expected.RightColumn[i])
		}
	}
}
