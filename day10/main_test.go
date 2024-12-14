package main

import "testing"

func TestTruc(t *testing.T) {
	puzzle := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	result := part1(puzzle)
	if result != 2 {
		t.Fatalf("Expected 2, received %d", result)
	}
}

func TestPart1(t *testing.T) {
	puzzle := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	result := part1(puzzle)
	if result != 36 {
		t.Fatalf("Expected 36, received %d", result)
	}
}

func TestPart2(t *testing.T) {
	puzzle := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	result := part2(puzzle)
	if result != 81 {
		t.Fatalf("Expected 81, received %d", result)
	}
}
