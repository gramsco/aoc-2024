package main

import (
	"slices"
)

type Graph struct {
	vertices []string
	edges    map[string][]string
}

func (g *Graph) append(a string, b string) {
	g.vertices = append(g.vertices, a, b)
	g.edges[a] = append(g.edges[a], b)
}

func (g Graph) filter(nodes []string) Graph {
	graph := Graph{
		vertices: []string{},
		edges:    make(map[string][]string),
	}
	for _, vertex := range g.vertices {
		if slices.Contains(nodes, vertex) {
			for _, edge := range g.edges[vertex] {
				graph.append(vertex, edge)
			}
		}
	}
	return graph
}

func (g *Graph) topologicalSort() []string {
	indegrees := map[string]int{}

	for _, v := range g.vertices {
		indegrees[v] = 0
	}

	for _, v := range g.edges {
		for _, neighbor := range v {
			indegrees[neighbor] += 1
		}
	}
	nodesWithNoIncomingEdges := []string{}

	for k, v := range indegrees {
		if v == 0 {
			nodesWithNoIncomingEdges = append(nodesWithNoIncomingEdges, k)
		}
	}

	sortedNodes := []string{}

	for {
		if len(nodesWithNoIncomingEdges) == 0 {
			break
		}
		node := nodesWithNoIncomingEdges[0]
		sortedNodes = append(sortedNodes, node)
		nodesWithNoIncomingEdges = nodesWithNoIncomingEdges[1:]

		for _, neighbor := range g.edges[node] {
			indegrees[neighbor]--
			if indegrees[neighbor] == 0 {
				nodesWithNoIncomingEdges = append(nodesWithNoIncomingEdges, neighbor)
			}
		}

	}

	return sortedNodes

}
