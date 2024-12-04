package main

import "testing"

func TestXxx(t *testing.T) {
	puzzle := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := solvePart1(
		puzzle)
	if result != 18 {
		t.Fatalf("expected 18, received %d", result)
	}
}
