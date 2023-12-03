package main

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key   int
	Data  string
	Edges []*Vertex
}

type Node struct {
	val     *Vertex
	visited bool
	next    *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(inValue *Vertex, visited bool) {
	node := &Node{val: inValue, visited: visited}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) Dequeue() *Vertex {
	returnValue := q.head.val
	q.head = q.head.next
	q.size--
	return returnValue
}

func (graph *Graph) AddVertex(key int, inputData string) {
	for _, val := range graph.Vertices {
		if key == val.Key {
			panic("\nKey already exists in graph!\n")
		}
	}
	graph.Vertices = append(graph.Vertices, &Vertex{Key: key, Data: inputData})
}

func (graph *Graph) PrintGraph() {
	for _, val := range graph.Vertices {
		fmt.Printf("\nVertex %v has data: %v and Edges:", val.Key, val.Data)
		for _, val := range val.Edges {
			fmt.Printf("%v", val.Key)
		}
	}
}

func (graph *Graph) getVertex(key int) *Vertex {
	for i, val := range graph.Vertices {
		if val.Key == key {
			return graph.Vertices[i]
		}
	}
	return nil
}

func (graph *Graph) AddEdge(fromVertexKey int, toVertexKey int) {
	fromVertex := graph.getVertex(fromVertexKey)
	toVertex := graph.getVertex(toVertexKey)

	fromVertex.Edges = append(fromVertex.Edges, toVertex)
}

func main() {
	g := &Graph{}
	g.AddVertex(1, "hello")
	g.AddVertex(2, "world")
	g.AddVertex(3, "Foo")
	g.AddEdge(1, 2)
	g.AddEdge(2, 1)
	g.AddEdge(1, 3)
	g.PrintGraph()
	q := &Queue{}
	for _, vert := range g.Vertices {
		q.Enqueue(vert, true)
	}
	fmt.Printf("\nSize %v\n", q.size)
	fmt.Printf("\nData at head %v", q.head.val.Edges)
	headEdges := q.head.val.Edges
	nextEdges := q.head.next.val.Edges
	for _, edge := range headEdges {
		fmt.Printf("\nEdge Key %v", edge.Key)
	}
	for _, edge := range nextEdges {
		fmt.Printf("\n Next Edge Key %v", edge.Key)
	}
}
