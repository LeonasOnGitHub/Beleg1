package graph

type Vertex struct {
	id    string
	edges []*Edge
}

type Edge struct {
	head   *Vertex
	tail   *Vertex
	length float64
}

type AdjacencyList struct {
	edges      []*Edge
	vertices   []*Vertex
	isDirected bool
}

func (x *AdjacencyList) addVertex(nodeId string) {
	v := Vertex{nodeId, []*Edge{}}
	p := &v
	x.vertices = append(x.vertices, p)
}

func (x *AdjacencyList) addEdge(nodeId1, nodeId2 string, length float64) {
	for _, v1 := range x.vertices {
		if nodeId1 == v1.id {
			for _, v2 := range x.vertices {
				if nodeId2 == v2.id {
					e := Edge{v1, v2, length}
					p := &e
					x.edges = append(x.edges, p)
				}
			}

		}
	}

}

func (x *AdjacencyList) DFS(nodeId, string) {
	var graph map[bool]Vertex
	for _, v1 := range x.vertices {
		graph[false] = v1
	}

}
