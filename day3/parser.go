package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Parser(s string) []Multiplication {
	var multiplications []Multiplication
	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		if currentChar == 'm' {
			substr := s[i:]
			if strings.HasPrefix(substr, "mul(") {
				rlIndex := strings.Index(substr, ")")
				multiplications = append(multiplications, getMultiplication(substr[:rlIndex+1]))
			}
		}

	}
	return multiplications
}

func ParserWithInstructions(s string) []Multiplication {
	enabled := true
	var multiplications []Multiplication
	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		if currentChar == 'd' {
			substr := s[i:]
			if strings.HasPrefix(substr, "don't()") {
				enabled = false
			}
			if strings.HasPrefix(substr, "do()") {
				enabled = true
			}
		}

		if currentChar == 'm' && enabled {
			substr := s[i:]
			if strings.HasPrefix(substr, "mul(") {
				rlIndex := strings.Index(substr, ")")
				multiplications = append(multiplications, getMultiplication(substr[:rlIndex+1]))
			}
			// TODO: jump to next
		}
	}
	return multiplications
}

func getMultiplication(s string) Multiplication {

	// mul(a,b) where 0 <= a <== 999 && 0 <= b <== 999
	r, err := regexp.Compile(`^mul\((\d{1,3}),(\d{1,3})\)$`)
	if err != nil {
		return Multiplication{}
	}
	matches := r.FindStringSubmatch(s)

	if len(matches) != 3 {
		return Multiplication{}
	}

	a, err := strconv.Atoi(matches[1])
	if err != nil {
		return Multiplication{}
	}
	b, err := strconv.Atoi(matches[2])
	if err != nil {
		return Multiplication{}
	}

	return Multiplication{a, b}
}
