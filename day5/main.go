package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RulesAndPages struct {
	rules   Graph
	updates [][]string
}

func part1(rp RulesAndPages) int {
	sum := 0
	for _, update := range rp.updates {

		graph := rp.rules.filter(update)
		rules := graph.topologicalSort()

		if isUpdateValid(rules, update) {
			half, err := strconv.Atoi(update[len(update)/2])
			if err != nil {
				panic(err)
			}
			sum += half
		}
	}
	return sum
}

func part2(rp RulesAndPages) int {
	sum := 0
	for _, update := range rp.updates {

		graph := rp.rules.filter(update)
		rules := graph.topologicalSort()

		if !isUpdateValid(rules, update) {
			sum += getFixedHalf(rules, update)
		}
	}
	return sum
}

func main() {
	puzzle := ParsePuzzle()

	fmt.Printf("Part1 = %d\n", part1(puzzle))

	fmt.Printf("Part2 = %d\n", part2(puzzle))

}

func ParsePuzzle() RulesAndPages {
	graph := Graph{vertices: []string{}, edges: make(map[string][]string)}
	updates := [][]string{}
	file, err := os.Open("day5/puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "|") {
			splitted := strings.Split(txt, "|")
			graph.append(splitted[0], splitted[1])
		}
		if strings.Contains(txt, ",") {
			splitted := strings.Split(txt, ",")
			updates = append(updates, splitted)
		}
	}

	return RulesAndPages{
		rules:   graph,
		updates: updates,
	}
}

func isUpdateValid(sortedRules []string, update []string) bool {
	actualRules := []string{}

	for _, rule := range sortedRules {
		if slices.Contains(update, rule) {
			actualRules = append(actualRules, rule)
		}
	}

	return strings.Join(actualRules, "") == strings.Join(update, "")
}

func getFixedHalf(sortedRules []string, update []string) int {

	actualRules := []string{}

	for _, rule := range sortedRules {
		if slices.Contains(update, rule) {
			actualRules = append(actualRules, rule)
		}
	}

	half, err := strconv.Atoi(actualRules[len(actualRules)/2])
	if err != nil {
		panic(err)
	}
	return half

}
