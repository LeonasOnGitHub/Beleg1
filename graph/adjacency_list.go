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
	//get vertex from nodeId1
	for _, v1 := range x.vertices {
		if nodeId1 == v1.id {
			//get vertex from nodeId2
			for _, v2 := range x.vertices {
				if nodeId2 == v2.id {
					e := Edge{v1, v2, length}
					p := &e
					x.edges = append(x.edges, p)
					if !x.isDirected {
						e := Edge{v2, v1, length}
						p := &e
						x.edges = append(x.edges, p)
					}
				}
			}

		}
	}

}

func (x *AdjacencyList) DFS(nodeId string) map[string]bool {
	var reachableVertices map[string]bool
	reachableVertices[nodeId] = true

	// get Vertex from nodeId
	for _, v := range x.vertices {
		if v.id == nodeId {
			//get all heads from vertex nodeId edges
			for _, e := range x.edges {
				if e.tail == v {
					//checks if edging vertex is in the map, if not adds it
					if _, exist := reachableVertices[e.head.id]; !exist {
						x.DFS(e.head.id)
					}
				}
			}
		}
	}
	return reachableVertices
}

func (x *AdjacencyList) BFS(nodeId string) map[string]int {
	var reachableVertices map[string]int
	reachableVertices[nodeId] = 0
	q := Queue{}
	// get Vertex from nodeId and adds it to the queue
	for _, v := range x.vertices {
		if v.id == nodeId {
			q.QueUp(v)
		}
	}
	var i = 0
	for len(q.vertices) > 0 {
		v := q.dequeue()
		for _, e := range x.edges {
			if e.tail == v {
				//checks if edging vertex is in the map, if not adds it
				if _, exist := reachableVertices[e.head.id]; !exist {
					reachableVertices[e.head.id] = i
					q.QueUp(e.head)
				}
			}
		}
	}

	return reachableVertices
}

func (x *AdjacencyList) NumVertices() int {
	return len(x.vertices)
}
func (x *AdjacencyList) NumEdges() int {
	return len(x.edges)
}
func (x *AdjacencyList) IsDirected() bool {
	return x.isDirected
}
