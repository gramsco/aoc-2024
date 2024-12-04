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

func TestPart2(t *testing.T) {
	puzzle := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
	result := solvePart2(
		puzzle)
	if result != 9 {
		t.Fatalf("expected 9, received %d", result)
	}
}
