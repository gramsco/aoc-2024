package main

type Coord struct {
	x, y int
}

func (coord Coord) outOfBounds(maxHeight int, maxWidth int) bool {
	if coord.x < 0 || coord.y < 0 {
		return true
	}
	if coord.y >= maxHeight {
		return true
	}
	if coord.x >= maxWidth {
		return true
	}
	return false
}

func (coord Coord) up() Coord {
	return Coord{coord.x, coord.y - 1}
}

func (coord Coord) right() Coord {
	return Coord{coord.x + 1, coord.y}
}

func (coord Coord) down() Coord {
	return Coord{coord.x, coord.y + 1}
}

func (coord Coord) left() Coord {
	return Coord{coord.x - 1, coord.y}
}

func (coord Coord) canGoLeft(m map[Coord]string) bool {
	up := m[coord.left()]
	return up != "#"
}

func (coord Coord) canGoDown(m map[Coord]string) bool {
	up := m[coord.down()]
	return up != "#"
}

func (coord Coord) canGoUp(m map[Coord]string) bool {
	up := m[coord.up()]
	return up != "#"
}

func (coord Coord) canGoRight(m map[Coord]string) bool {
	up := m[coord.right()]
	return up != "#"
}
