package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	res := parse("2333133121414131402")
	db := Debug(res)
	if db != "00...111...2...333.44.5555.6666.777.888899" {
		t.Fatalf("Received %s", db)
	}

}

// func TestRegroup(t *testing.T) {
// 	result := Regroup(Expand(parse("2333133121414131402")))
// 	if result != "0099811188827773336446555566.............." {
// 		t.Fatalf("Received %s", result)
// 	}

// }

// func TestExpand(t *testing.T) {
// 	b := Expand([]Block{{51, 2, "file"}})
// 	if b != "5151" {
// 		t.Fatalf("boom -> %s\n", b)
// 	}
// 	c := Expand([]Block{{1, 2, "file"}, {-1, 1, "empty"}, {2, 1, "file"}})
// 	if c != "112." {
// 		t.Fatalf("boom -> %s\n", c)
// 	}

// }

func TestPart1(t *testing.T) {
	result := part1("2333133121414131402")
	if result != 1928 {
		t.Fatalf("Expected 1928, got %d", result)
	}

	// result2 := part1("12345")
	// if result != 0 {
	// 	t.Fatalf("Result %d", result2)
	// }
}
