package main

import (
	"testing"
)

func TestSolve1(t *testing.T) {

	puzzle := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := SolvePart1(puzzle)
	if SolvePart1(puzzle) != 2 {
		t.Fatalf("Received %d, expected 2", result)
	}
}

func TestSolve2(t *testing.T) {
	puzzle := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := SolvePart2(puzzle)
	if SolvePart2(puzzle) != 4 {
		t.Fatalf("Received %d, expected 4", result)
	}
}

func TestCasesPart1(t *testing.T) {

	// Safe
	for i, test := range [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 6, 7, 9},
	} {
		if !CheckReportSafetyWithDampener(test) {
			t.Fatalf("Case %d should be safe!\n", i+1)
		}
	}

	// Unsafe
	for i, test := range [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
	} {
		if CheckReportSafety(test) {
			t.Fatalf("Unsafe case %d!\n", i+1)
		}
	}

}

func TestCasesPart2(t *testing.T) {

	// Safe
	for i, test := range [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	} {
		if !CheckReportSafetyWithDampener(test) {
			t.Fatalf("Case %d should be safe!\n", i+1)
		}
	}

	// Unsafe
	for i, test := range [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
	} {
		if CheckReportSafetyWithDampener(test) {
			t.Fatalf("Case %d should be unsafe!\n", i+1)
		}
	}

}
