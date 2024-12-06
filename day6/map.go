package main

import (
	"fmt"
	"log"
)

func makeMap(lines []string) (map[Coord]string, Coord) {

	guard := Coord{}
	m := make(map[Coord]string)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			line := lines[y]
			value := string(line[x])
			if value == "^" {
				guard = Coord{x, y}

			}
			m[Coord{x, y}] = value
		}
	}
	return m, guard
}

func isLoopyObstacle(lines []string, obstacle Coord) bool {

	m, guard := makeMap(lines)

	if guard.x == obstacle.x && guard.y == obstacle.y {
		return false
	}

	m[obstacle] = "#"

	positions := map[Coord]int{}

	for {

		positions[guard]++

		if guard.outOfBounds(len(lines), len(lines[0])) {
			return false
		}

		if positions[guard] > 4 {
			return true
		}

		cursor := m[guard]

		switch cursor {
		case ".":
			fmt.Printf("Guard position x : %d, y : %d\n", guard.x, guard.y)
			log.Fatalf("Wrong cursor : %s", cursor)

		case "^":
			if guard.canGoUp(m) {
				m[guard] = "X"
				guard = guard.up()
				m[guard] = "^"
				continue
			} else {
				m[guard] = ">"
				continue
			}
		case ">":
			if guard.canGoRight(m) {
				m[guard] = "X"
				guard = guard.right()
				m[guard] = ">"
				continue
			} else {
				m[guard] = "v"
				continue
			}
		case "v":
			if guard.canGoDown(m) {
				m[guard] = "X"
				guard = guard.down()
				m[guard] = "v"
				continue
			} else {
				m[guard] = "<"
				continue
			}
		case "<":
			if guard.canGoLeft(m) {
				m[guard] = "X"
				guard = guard.left()
				m[guard] = "<"
				continue
			} else {
				m[guard] = "^"
				continue
			}
		}
	}

}

func getGuardPath(lines []string) []Coord {

	m, guard := makeMap(lines)

	for {

		if guard.outOfBounds(len(lines), len(lines[0])) {
			break
		}

		cursor := m[guard]

		switch cursor {
		case ".":
			fmt.Printf("Guard position x : %d, y : %d\n", guard.x, guard.y)
			log.Fatalf("Wrong cursor : %s", cursor)

		case "^":
			if guard.canGoUp(m) {
				m[guard] = "X"
				guard = guard.up()
				m[guard] = "^"
				continue
			} else {
				m[guard] = ">"
				continue
			}
		case ">":
			if guard.canGoRight(m) {
				m[guard] = "X"
				guard = guard.right()
				m[guard] = ">"
				continue
			} else {
				m[guard] = "v"
				continue
			}
		case "v":
			if guard.canGoDown(m) {
				m[guard] = "X"
				guard = guard.down()
				m[guard] = "v"
				continue
			} else {
				m[guard] = "<"
				continue
			}
		case "<":
			if guard.canGoLeft(m) {
				m[guard] = "X"
				guard = guard.left()
				m[guard] = "<"
				continue
			} else {
				m[guard] = "^"
				continue
			}
		}
	}

	coords := []Coord{}

	for coord, v := range m {
		if v == "X" {
			coords = append(coords, coord)
		}
	}

	return coords

}
