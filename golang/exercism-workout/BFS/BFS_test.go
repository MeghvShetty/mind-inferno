package main

import (
	"bytes"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][2]int
		start    int
		expected string
	}{
		{
			name: "Simple path graph",
			edges: [][2]int{
				{0, 1},
				{1, 2},
				{2, 3},
			},
			start:    0,
			expected: "0 1 2 3 ",
		},
		{
			name: "Example graph from main",
			edges: [][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{2, 4},
				{3, 5},
				{4, 5},
			},
			start:    0,
			expected: "0 1 2 3 4 5 ",
		},
		{
			name: "Disconnected graph",
			edges: [][2]int{
				{0, 1},
				{1, 2},
				{3, 4},
				{4, 5},
			},
			start:    0,
			expected: "0 1 2 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new graph for each test case
			g := NewGraph()
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			// Create a buffer to capture output
			var buf bytes.Buffer

			// Run BFS
			BFS(g, tt.start, &buf)

			// Compare the output
			if got := buf.String(); got != tt.expected {
				t.Errorf("BFS() = %v, want %v", got, tt.expected)
			}
		})
	}
}
