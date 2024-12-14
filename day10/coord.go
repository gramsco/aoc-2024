package main

import "strconv"

type Coord struct {
	y int
	x int
}

func (c *Coord) isNil() bool {
	return c.x == -1 || c.y == -1
}

func NilCoord() Coord {
	return Coord{-1, -1}
}

func NewCoord(s string, s2 string) Coord {
	y, _ := strconv.Atoi(s)
	x, _ := strconv.Atoi(s2)
	return Coord{y, x}
}

func CoordOrNil(puzzle []string, coord Coord, target int) Coord {
	value, error := strconv.Atoi(string(puzzle[coord.y][coord.x]))
	if value != target {
		return NilCoord()
	}
	if error != nil {
		return NilCoord()
	}
	return Coord{coord.y, coord.x}
}

func Top(puzzle []string, coord Coord, target int) Coord {
	if coord.y <= 0 {
		return NilCoord()
	}
	return CoordOrNil(puzzle, Coord{coord.y - 1, coord.x}, target)
}

func Down(puzzle []string, coord Coord, target int) Coord {
	if coord.y >= len(puzzle)-1 {
		return NilCoord()
	}
	return CoordOrNil(puzzle, Coord{coord.y + 1, coord.x}, target)
}

func Left(puzzle []string, coord Coord, target int) Coord {
	if coord.x <= 0 {
		return NilCoord()
	}
	return CoordOrNil(puzzle, Coord{coord.y, coord.x - 1}, target)
}

func Right(puzzle []string, coord Coord, target int) Coord {
	if coord.x >= len(puzzle[0])-1 {
		return NilCoord()
	}
	return CoordOrNil(puzzle, Coord{coord.y, coord.x + 1}, target)
}
