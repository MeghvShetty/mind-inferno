// This is a sample for DFS (depth-first search)

package main

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func DFS(g *Graph, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	neighbors := g.adjList[start]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			DFS(g, neighbor, visited)
		}
	}
}

func main() {
	g := NewGraph()

	// Example graph
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	startNode := 0
	visited := make(map[int]bool)

	fmt.Println("Depth-First Traversal (DFS):")
	DFS(g, startNode, visited)
}
