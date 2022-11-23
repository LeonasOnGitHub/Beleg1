package tests

import (
	"Beleg1/graph"
	"fmt"
	"testing"
)

var vertices = make([]*graph.Vertex, 0)
var edges = make([]*graph.Edge, 0)
var al = graph.AdjacencyList{edges, vertices, true}

func TestMain(m *testing.M) {

	al.AddVertex("A")
	al.AddVertex("B")
	al.AddVertex("C")
	al.AddVertex("D")
	al.AddVertex("E")

	al.AddEdge("A", "B", 1)
	al.AddEdge("A", "C", 1)
	al.AddEdge("B", "C", 1)
	al.AddEdge("B", "D", 1)
	al.AddEdge("C", "D", 1)

	m.Run()
}

func TestVertices(t *testing.T) {
	for i := 0; i < al.NumVertices(); i++ {
		fmt.Print(al.Vertices[i])
	}
	fmt.Println()
}
func TestDFS(t *testing.T) {
	resultDFS := al.DFS("A")
	var dummyArray [5]string
	dummyArray[0] = "A"
	dummyArray[1] = "B"
	dummyArray[2] = "C"
	dummyArray[3] = "D"
	dummyArray[4] = "E"
	for i := 0; i < len(dummyArray); i++ {
		_, ok := resultDFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
	}
}
