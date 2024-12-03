package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Multiplication struct {
	a int
	b int
}

func main() {
	lines := newPuzzleFromInput()
	sum := parseAndCompute(lines, Parser)
	fmt.Printf("Part1 : %d\n", sum)

	sum2 := parseAndCompute(lines, ParserWithInstructions)
	fmt.Printf("Part2 : %d\n", sum2)
}

func parseAndCompute(p string, parser func(s string) []Multiplication) int {
	multiplications := parser(p)
	result := 0
	for _, multiplication := range multiplications {
		result += multiplication.a * multiplication.b
	}
	return result
}

func newPuzzleFromInput() string {

	file, err := os.Open("day3/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := ""
	for scanner.Scan() {
		lines += scanner.Text()
	}
	return lines

}
