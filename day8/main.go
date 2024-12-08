package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	puzzle := readPuzzle()
	fmt.Printf("Result is %d\n", part1(puzzle))
	fmt.Printf("Result part 2 is %d\n", part2(puzzle))
}

func readPuzzle() string {
	file, err := os.ReadFile("day8/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func part1(s string) int {
	m := make(map[Vector]bool)
	p, height, width := parsePuzzle(s)

	for _, v := range p {
		l := computeFrequencyLocation(v, height, width)
		for _, vector := range l {
			if vector.isInBounds(height, width) {
				m[vector] = true
			}
		}
	}

	return len(m)
}

func part2(s string) int {
	m := make(map[Vector]bool)
	p, height, width := parsePuzzle(s)

	for _, v := range p {
		l := computeFrequencyLocationPart2(v, height, width)
		for _, vector := range l {
			m[vector] = true
		}
	}

	return len(m)
}

func computeFrequencyLocationPart2(locations []Vector, height int, width int) []Vector {
	res := []Vector{}
	m := make(map[Vector]string)

	for _, l1 := range locations {
		for _, l2 := range locations {
			if l1 == l2 {
				continue
			}

			vec := l2.minus(l1)
			x := l1

			for {
				x = x.minus(vec)
				if x.isInBounds(height, width) {
					_, ok := m[x]
					if !ok {
						m[x] = "ok"
					}
				} else {
					break
				}

			}

			y := l2
			for {
				y = y.minus(vec)
				if y.isInBounds(height, width) {
					_, ok := m[y]
					if !ok {
						m[y] = "ok"
					}
				} else {
					break
				}

			}
		}
	}

	for k := range m {
		res = append(res, k)
	}
	return res
}

func computeFrequencyLocation(locations []Vector, height int, width int) []Vector {
	res := []Vector{}
	m := make(map[Vector]bool)

	for _, l1 := range locations {
		for _, l2 := range locations {
			if l1 == l2 {
				continue
			}
			vec := l2.minus(l1)
			x := l1.minus(vec)
			y := l2.plus(vec)

			if x.isInBounds(height, width) {
				if _, ok := m[x]; !ok {
					m[x] = true
				}
			}

			if y.isInBounds(height, width) {
				if _, ok := m[y]; !ok {
					m[y] = true
				}
			}

		}
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func parsePuzzle(s string) (map[rune][]Vector, int, int) {

	result := make(map[rune][]Vector)
	lines := strings.Split(s, "\n")

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				r := result[c]
				result[c] = append(r, Vector{y, x})
			}
		}
	}
	return result, len(lines), len(lines[0])
}
