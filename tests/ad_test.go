package tests

import (
	"Beleg1/graph"
	"fmt"
	"testing"
)

var vertices = make([]*graph.Vertex, 0)
var edges = make([]*graph.Edge, 0)
var alD = graph.AdjacencyList{Edges: edges, Vertices: vertices, IsDirected: true}
var alUD = graph.AdjacencyList{Edges: edges, Vertices: vertices, IsDirected: false}

func TestMain(m *testing.M) {

	alD.AddVertex("A")
	alD.AddVertex("B")
	alD.AddVertex("C")
	alD.AddVertex("D")
	alD.AddVertex("E")
	alD.AddEdge("A", "B", 1)
	alD.AddEdge("A", "C", 1)
	alD.AddEdge("B", "C", 1)
	alD.AddEdge("B", "D", 1)
	alD.AddEdge("C", "D", 1)

	alUD.AddVertex("A")
	alUD.AddVertex("B")
	alUD.AddVertex("C")
	alUD.AddVertex("D")
	alD.AddEdge("A", "B", 1)
	alD.AddEdge("A", "C", 1)
	alD.AddEdge("A", "D", 1)

	m.Run()
}

func TestVertices(t *testing.T) {
	for i := 0; i < alD.NumVertices(); i++ {
		fmt.Print(alD.Vertices[i])
	}
	fmt.Println()
}
func TestEdges(t *testing.T) {
	for i := 0; i < alD.NumEdges(); i++ {
		v1 := *alD.Edges[i].Tail
		v2 := *alD.Edges[i].Head
		fmt.Println("Tail: ", v1, "Head: ", v2)
	}
	fmt.Println()
}

func TestDFSDirected(t *testing.T) {
	resultDFS := alD.DFS("A")
	var dummyArray [4]string
	dummyArray[0] = "A"
	dummyArray[1] = "B"
	dummyArray[2] = "C"
	dummyArray[3] = "D"
	if len(dummyArray) != len(resultDFS) {
		t.Error("The dummyArray is not the same size as the resultArray")
	}
	for i := 0; i < len(dummyArray); i++ {
		_, ok := resultDFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
	}
}
func TestDFSUndirected(t *testing.T) {
	resultDFS := alUD.DFS("A")
	var dummyArray [4]string
	dummyArray[0] = "A"
	dummyArray[1] = "B"
	dummyArray[2] = "C"
	dummyArray[3] = "D"
	if len(dummyArray) != len(resultDFS) {
		t.Error("The dummyArray is not the same size as the resultArray")
	}
	for i := 0; i < len(dummyArray); i++ {
		_, ok := resultDFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
	}
}
func TestBFSUndirected(t *testing.T) {
	resultDFS := alUD.BFS("A")
	var dummyArray [4]string
	dummyArray[0] = "A"
	dummyArray[1] = "B"
	dummyArray[2] = "C"
	dummyArray[3] = "D"
	if len(dummyArray) != len(resultDFS) {
		t.Error("The dummyArray is not the same size as the resultArray")
	}
	for i := 0; i < len(dummyArray); i++ {
		_, ok := resultDFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
	}
}
