package main

import (
	"fmt"
	"log"
)

// Graph Struct (as adjacency list)
type Graph struct {
	vertices int     // N vertices
	edges    [][]int // Adjacency list for edges
}

// Adds a directed edge from the vertex [vertexFrom] to the vertex [vertexTo]
func (g *Graph) AddEdges(vertexFrom int, vertexTo int) {
	g.edges[vertexFrom] = append(g.edges[vertexFrom], vertexTo)
}

func GraphConstructor(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]int, vertices),
	}
}

// Recursive utility for traversing the graph from a starting vertex
// Then, keep track of the visited nodes with a slice of bools.
// Visit all adjacent nodes recursively.
func (g *Graph) utility(start int, visited []bool) {
	visited[start] = true // Marks the initial node as visited so we don't come back to it.
	log.Printf("On Vertex %v \n", start)

	for _, i := range g.edges[start] {
		if !visited[i] { // If a vertex adjacent to the one passed is not visited.
			g.utility(i, visited) // Recurse
		}
	}
}

// Traverse
func (g *Graph) DFS(startVertex int) {
	visited := make([]bool, g.vertices)
	g.utility(startVertex, visited)
}

func main() {
	g := GraphConstructor(9)

	g.AddEdges(0, 1)
	g.AddEdges(0, 2)
	g.AddEdges(1, 2)
	g.AddEdges(2, 0)
	g.AddEdges(3, 1)
	g.AddEdges(4, 1)
	g.AddEdges(5, 1)
	g.AddEdges(6, 1)
	g.AddEdges(7, 1)
	g.AddEdges(8, 1)
	g.AddEdges(8, 2)
	g.AddEdges(8, 3)
	g.AddEdges(8, 4)
	g.AddEdges(8, 5)
	g.AddEdges(8, 6)
	g.AddEdges(8, 7)

	fmt.Println("Attempting to Traverse using DFS: ")
	g.DFS(8)
}
