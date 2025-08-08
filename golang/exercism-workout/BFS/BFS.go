package main

import (
	"fmt"
	"io"
	"os"
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

func BFS(g *Graph, start int, writer io.Writer) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Fprintf(writer, "%d ", node)

		neighbors := g.adjList[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
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

	fmt.Println("Breadth-First Traversal (BFS):")
	BFS(g, startNode, os.Stdout)
}
