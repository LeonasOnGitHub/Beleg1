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
	alD.AddEdge("A", "B", 0.5)
	alD.AddEdge("A", "C", 1)
	alD.AddEdge("B", "C", 0.4)
	alD.AddEdge("B", "D", 1)
	alD.AddEdge("C", "D", 0.3)

	alUD.AddVertex("A")
	alUD.AddVertex("B")
	alUD.AddVertex("C")
	alUD.AddVertex("D")
	alUD.AddEdge("A", "B", 1)
	alUD.AddEdge("A", "C", 1)
	alUD.AddEdge("A", "D", 1)

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
		t.Error("The dummyArray is not the same size as the resultMap")
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
		t.Error("The dummyArray is not the same size as the resultMap")
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
	var dummyArraylenght [4]int
	dummyArraylenght[0] = 0
	dummyArraylenght[1] = 1
	dummyArraylenght[2] = 1
	dummyArraylenght[3] = 1
	if len(dummyArray) != len(resultDFS) {
		t.Error("The dummyArray is not the same size as the resultMap")
	}
	for i := 0; i < len(dummyArray); i++ {
		v, ok := resultDFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
		if v != dummyArraylenght[i] {
			t.Error("The Vertex should have the length ", dummyArraylenght[i], " but instead has the length ", v)
		}
	}
}
func TestBFSDirected(t *testing.T) {
	resultBFS := alD.BFS("A")
	var dummyArray [4]string
	dummyArray[0] = "A"
	dummyArray[1] = "B"
	dummyArray[2] = "C"
	dummyArray[3] = "D"
	var dummyArrayLenght [4]int
	dummyArrayLenght[0] = 0
	dummyArrayLenght[1] = 1
	dummyArrayLenght[2] = 1
	dummyArrayLenght[3] = 2
	if len(dummyArray) != len(resultBFS) {
		t.Error("The dummyArray is not the same size as the resultMap")
	}
	for i := 0; i < len(dummyArray); i++ {
		v, ok := resultBFS[dummyArray[i]]
		if !ok {
			t.Error("The Vertex ", dummyArray[i], " should be in map but isn´t")
		}
		if v != dummyArrayLenght[i] {
			t.Error("The Vertex should have the length ", dummyArrayLenght[i], " but instead has the length ", v)
		}
	}
}
func TestTopoSort(t *testing.T) {
	resultTopo := alD.TopoSort()
	var dummyMapDist = make(map[string]int)
	dummyMapDist["A"] = 1
	dummyMapDist["B"] = 2
	dummyMapDist["C"] = 3
	dummyMapDist["D"] = 4

	for key, value := range resultTopo {

		if value != dummyMapDist[key] {
			t.Error("The Vertex ", key, " should have the order", dummyMapDist[key], " but has ", value)
		}
	}
}

func TestDijkstra(t *testing.T) {
	resultDijk := alD.Dijkstra("A")
	var dummyMapDist = make(map[string]float64)
	dummyMapDist["A"] = 0
	dummyMapDist["B"] = 0.5
	dummyMapDist["C"] = 0.9
	dummyMapDist["D"] = 1.2

	for key, value := range resultDijk {
		if value != dummyMapDist[key] {
			t.Error("The Vertex ", key, " should have the length", dummyMapDist[key], " but has ", value)
		}
	}
}
