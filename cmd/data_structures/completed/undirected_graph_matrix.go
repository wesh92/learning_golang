package main

import "fmt"

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Value int
	Edges map[int]*Edge
}

type Edge struct {
	Weight int
	Vertex *Vertex
}

func (graph *Graph) AddVertex(key int, inValue int) {
	graph.Vertices[key] = &Vertex{Value: inValue, Edges: map[int]*Edge{}}
}

func (graph *Graph) KeyExists(key int) bool {
	if _, ok := graph.Vertices[key]; ok {
		return ok
	}
	return false
}

func (graph *Graph) AddEdge(sourceKey int, destinationKey int, weight int) {
	if ok := graph.KeyExists(sourceKey); !ok {
		return
	}
	if ok := graph.KeyExists(destinationKey); !ok {
		return
	}

	graph.Vertices[sourceKey].Edges[destinationKey] = &Edge{Weight: weight, Vertex: graph.Vertices[destinationKey]}
}

func (graph *Graph) Neighbors(sourceKey int) []int {
	resultList := []int{}

	if _, ok := graph.Vertices[sourceKey]; !ok {
		return resultList
	}

	for _, edge := range graph.Vertices[sourceKey].Edges {
		resultList = append(resultList, edge.Vertex.Value)
	}

	return resultList
}

func main() {
	g := &Graph{Vertices: map[int]*Vertex{}}
	g.AddVertex(1, 32)
	g.AddVertex(2, 64)
	g.AddVertex(4, 128)
	g.AddEdge(1, 2, 5)
	g.AddEdge(1, 4, 6)
	fmt.Printf("Exists Check: %v\n", g.KeyExists(2))
	fmt.Printf("Value: %p %v\n", &g, g.Vertices[1].Value)
	fmt.Printf("Value: %p %v\n", &g, g.Vertices[2].Value)
	fmt.Printf("Value: %p %v\n", &g, g.Vertices[1].Edges[2].Vertex.Value)
	for _, neighbor := range g.Neighbors(1) {
		fmt.Println(neighbor)
	}
}
