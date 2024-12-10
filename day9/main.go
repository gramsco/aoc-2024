package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	res := part1(newPuzzleFromInput())
	fmt.Println(res)
}

func newPuzzleFromInput() string {

	file, err := os.Open("day9/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzle := ""
	for scanner.Scan() {
		puzzle += scanner.Text()
	}

	return puzzle
}

type Block struct {
	id   int
	size int
	kind string
}

func parse(s string) []Block {
	results := []Block{}

	id := -1

	for i, result := range s {
		size, err := strconv.Atoi(string(result))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			id++
			results = append(results, Block{id, size, "file"})
		} else {
			// if size == 0 {
			// 	continue
			// }
			results = append(results, Block{-1, size, "empty"})
		}

	}
	return results
}

func Insert[T any](slice []T, index int, element T) []T {
	// Ensure the index is valid
	if index < 0 || index > len(slice) {
		panic("index out of range")
	}

	// Insert the element
	slice = append(slice[:index], append([]T{element}, slice[index:]...)...)
	return slice
}

func Debug(blocks []Block) string {
	s := ""
	for _, b := range blocks {
		for range b.size {
			if b.kind == "file" {
				s += strconv.Itoa(b.id)
			} else {
				s += "."
			}
		}
	}
	return s
}

func Expand(blocks []Block) string {
	s := ""

	totalToAdd := 0

	code := 0

	for _, b := range blocks {
		if b.kind == "file" {
			totalToAdd += len(strconv.Itoa(b.id)) * b.size
		}
	}

	rightCursor := len(blocks) - 1
	consumed := 0
	copy := []Block{}
	totalRemoved := 0

	for _, b := range blocks {

		if len(s) >= totalToAdd {
			break
		}

		if b.kind == "file" {
			code += len(copy) * b.id * b.size
			copy = append(copy, b)
			continue
		}

		if b.kind == "empty" {
			to_fill := b.size
			filled := 0

			for {
				if filled == to_fill {
					break
				}
				if rightCursor == 0 {
					break
				}
				next := blocks[rightCursor]
				if next.kind == "empty" {
					rightCursor--
				} else {
					copy = append(copy, Block{id: next.id, size: 1, kind: "file"})
					consumed++
					filled++
					totalRemoved++
					if consumed == next.size {
						consumed = 0
						rightCursor--
						continue
					}
				}

			}

		}
	}

	fmt.Println(Debug(blocks))
	fmt.Println(Debug(copy))
	c := Debug(copy)
	return c[:totalToAdd]

}

func part1(s string) int {
	parsed := parse(s)
	res := Expand(parsed)

	fmt.Println(res)

	sum := 0

	for i, c := range res {
		b, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		sum += i * b
	}
	return sum
}
