package main

type Vector struct {
	y int
	x int
}

func (v Vector) minus(l2 Vector) Vector {
	return Vector{v.y - l2.y, v.x - l2.x}
}

func (v Vector) plus(l2 Vector) Vector {
	return Vector{v.y + l2.y, v.x + l2.x}
}

func (vector Vector) isInBounds(height int, width int) bool {
	if vector.x < 0 {
		return false
	}
	if vector.x >= width {
		return false
	}
	if vector.y < 0 {
		return false
	}
	if vector.y >= height {
		return false
	}
	return true
}
