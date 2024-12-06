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

func main() {
	rp := ParsePuzzle()

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

	fmt.Println(sum)
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
