package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result part1 : %d\n", part1())
}

func part1() int {
	file, err := os.Open("day7/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		txt := scanner.Text()
		parsed := parseStr(txt)
		count += ResultOrZero(parsed[0], parsed[1:]...)

	}

	return count
}

func calculate(n []int, operations []string) int {
	result := n[0]
	for i := 0; i < len(operations); i++ {
		operation := operations[i]
		to_add := n[i+1]
		if operation == "add" {
			result = result + to_add
		}
		if operation == "multiply" {
			result = result * to_add
		}

	}
	return result
}

// Ok that one is from chatGPT -_-
func binaryString(num, width int) string {
	binary := strconv.FormatInt(int64(num), 2)
	// Pad with leading zeros if necessary
	if len(binary) < width {
		padding := width - len(binary)
		return fmt.Sprintf("%0*s", padding+len(binary), binary)
	}
	// Trim to width if binary length exceeds width
	return binary[len(binary)-width:]
}

func ResultOrZero(total int, elements ...int) int {
	for operation := range int(math.Pow(2.0, float64(len(elements)-1))) {
		list := ExtractOperation(operation, len(elements)-1)
		r := calculate(elements, list)
		if total == r {
			return r
		}
	}
	return 0
}

func parseStr(s string) []int {
	result := []int{}

	splitted := strings.Split(strings.Replace(s, ":", "", 1), " ")

	for _, el := range splitted {
		parsed, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		result = append(result, parsed)
	}

	return result
}

func ExtractOperation(binary int, length int) []string {

	s := binaryString(binary, length)

	result := []string{}
	for _, c := range s {
		if c == '0' {
			result = append(result, "add")
		} else {
			result = append(result, "multiply")
		}
	}
	return result
}
