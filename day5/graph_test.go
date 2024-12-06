package main

import (
	"slices"
	"testing"
)

func TestAppendEdge(t *testing.T) {
	graph := Graph{
		vertices: []string{},
		edges:    make(map[string][]string),
	}
	graph.append("1", "2")

	if !slices.Contains(graph.vertices, "2") {
		t.Fatalf("2 should be in graph")
	}

	if !slices.Contains(graph.edges["1"], "2") {
		t.Fatalf("2 should be in graph")
	}
}

func TestSort(t *testing.T) {
	graph := Graph{
		vertices: []string{},
		edges:    make(map[string][]string),
	}

	graph.append("47", "53")
	graph.append("97", "13")
	graph.append("97", "61")
	graph.append("97", "47")
	graph.append("75", "29")
	graph.append("61", "13")
	graph.append("75", "53")
	graph.append("29", "13")
	graph.append("97", "29")
	graph.append("53", "29")
	graph.append("61", "53")
	graph.append("97", "53")
	graph.append("61", "29")
	graph.append("47", "13")
	graph.append("75", "47")
	graph.append("97", "75")
	graph.append("47", "61")
	graph.append("75", "61")
	graph.append("47", "29")
	graph.append("75", "13")
	graph.append("53", "13")

	expected := []string{"97", "75", "47", "61", "53", "29", "13"}

	result := graph.topologicalSort()
	for i := range result {
		if expected[i] != result[i] {
			t.Fatalf("Expected %s, received %s", expected[i], result[i])
		}
	}
}

func TestValidUpdate(t *testing.T) {
	rules := []string{"97", "75", "47", "61", "53", "29", "13"}

	// VALID
	if !isUpdateValid(rules, []string{"75", "47", "61", "53", "29"}) {
		t.Fatalf("Should be valid")
	}

	if !isUpdateValid(rules, []string{"97", "61", "53", "29", "13"}) {
		t.Fatalf("Should be valid")
	}

	if !isUpdateValid(rules, []string{"75", "29", "13"}) {
		t.Fatalf("Should be valid")
	}

	// INVALID
	if isUpdateValid(rules, []string{"75", "97", "47", "61", "53"}) {
		t.Fatalf("Should be valid")
	}

	if isUpdateValid(rules, []string{"61", "13", "29"}) {
		t.Fatalf("Should be valid")
	}

	if isUpdateValid(rules, []string{"97", "13", "75", "29", "47"}) {
		t.Fatalf("Should be valid")
	}

}

func TestFixInvalid(t *testing.T) {
	rules := []string{"97", "75", "47", "61", "53", "29", "13"}

	result := getFixedHalf(rules, []string{"75", "97", "47", "61", "53"})

	if result != 47 {
		t.Fatalf("Expected %d, received %d", 47, result)
	}

}
