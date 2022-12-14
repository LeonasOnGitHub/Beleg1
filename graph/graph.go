// Package graph implements an adjacency list for a graph representation
// and some selected algorithms on such graphs.
package graph

// Interface of a Graph
// - each node should have an unique nodeId (externally set) of type string
type Graph interface {

	// add an vertex. If the vertex already exists in the graph: do nothing
	AddVertex(nodeId string) //, data interface{})

	//HasVertex(nodeId string) bool
	AddEdge(nodeId1, nodeId2 string, length float64) //, data interface{})

	// get some generic meta data of the graph:
	NumVertices() int // number of Vertices
	NumEdges() int    // number of Edges
	IsDirected() bool // is a directed or undirected graph?

	// some basic algorithms:

	// the returned map maps for each reachable
	// node the distance in layers from given nodeId to each node
	BFS(nodeId string) map[string]int

	// maps the reachable node to true
	DFS(nodeId string) map[string]bool
}

type DirectedGraph interface {
	Graph
	Dijkstra(id string) map[string]float64
	//BellmannFord(Id string) ([]float64, bool)
	//FloydWarshall() ([][]float64, bool) // no negative cycles
}

// directed acyclic graph
type DAG interface {
	DirectedGraph
	// algorithms
	// TopoSort returns an map from nodeId to topological order
	TopoSort() map[string]int
}

type UnDirectedGraph interface {
	Graph
	//Kruskal(graph *AdjacencyList) (*AdjacencyList, float64)
}
