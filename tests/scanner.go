package tests

import (
	"Beleg1/graph"
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func initGraph9(filename string, t *testing.T) graph.AdjacencyList {

	vertices := make([]*graph.Vertex, 0)
	edges := make([]*graph.Edge, 0)
	graph := graph.AdjacencyList{Edges: edges, Vertices: vertices, IsDirected: true}

	file, _ := os.Open(filename) //
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",") // f[0]:id2 ,,
			var length float64         //edgeLength
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l //edgeLength{float64: l}
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddEdge(id1, f[0], length)
		}

	}
	return graph
}
