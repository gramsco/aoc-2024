package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result part1 : %d\n", scanPuzzle(false))
	fmt.Printf("Result part2 : %d\n", scanPuzzle(true))
}

func scanPuzzle(part2 bool) int {
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
		count += reduceCompute(parsed[0], parsed[1:], part2)
	}
	return count
}

func concat(a, b int) int {
	res, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(err)
	}
	return res
}

func reduceCompute(total int, elements []int, part2 bool) int {
	if len(elements) == 1 {
		if elements[0] == total {
			return total
		}
		return 0
	}
	if n := reduceCompute(total, append([]int{elements[0] + elements[1]}, elements[2:]...), part2); n != 0 {
		return n
	}
	if n := reduceCompute(total, append([]int{elements[0] * elements[1]}, elements[2:]...), part2); n != 0 {
		return n
	}

	if part2 {
		if n := reduceCompute(total, append([]int{concat(elements[0], elements[1])}, elements[2:]...), part2); n != 0 {
			return n
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
