package graph

type Vertex struct {
	Id    string
	Edges []*Edge
}

type Edge struct {
	Head   *Vertex
	Tail   *Vertex
	Length float64
}

type AdjacencyList struct {
	Edges      []*Edge
	Vertices   []*Vertex
	IsDirected bool
}

func (x *AdjacencyList) AddVertex(nodeId string) {
	v := Vertex{nodeId, []*Edge{}}
	p := &v
	x.Vertices = append(x.Vertices, p)
}

func (x *AdjacencyList) AddEdge(nodeId1, nodeId2 string, length float64) {
	//get vertex from nodeId1
	for _, v1 := range x.Vertices {
		if nodeId1 == v1.Id {
			//get vertex from nodeId2
			for _, v2 := range x.Vertices {
				if nodeId2 == v2.Id {
					e := Edge{v1, v2, length}
					p := &e
					x.Edges = append(x.Edges, p)
					if !x.IsDirected {
						e := Edge{v2, v1, length}
						p := &e
						x.Edges = append(x.Edges, p)
					}
				}
			}

		}
	}

}

func (x *AdjacencyList) DFS(nodeId string) map[string]bool {
	reachableVertices := make(map[string]bool)
	reachableVertices[nodeId] = true

	// get Vertex from nodeId
	for _, v := range x.Vertices {
		if v.Id == nodeId {
			//get all heads from vertex nodeId Edges
			for _, e := range x.Edges {
				if e.Tail == v {
					//checks if edging vertex is in the map, if not adds it
					if _, exist := reachableVertices[e.Head.Id]; !exist {
						x.DFS(e.Head.Id)
					}
				}
			}
		}
	}
	return reachableVertices
}

func (x *AdjacencyList) BFS(nodeId string) map[string]int {
	reachableVertices := make(map[string]int)
	reachableVertices[nodeId] = 0
	q := Queue{}
	// get Vertex from nodeId and adds it to the queue
	for _, v := range x.Vertices {
		if v.Id == nodeId {
			q.QueUp(v)
		}
	}
	for len(q.vertices) > 0 {
		v := q.dequeue()
		for _, e := range x.Edges {
			if e.Tail == v {
				//checks if edging vertex is in the map, if not adds it
				if _, exist := reachableVertices[e.Head.Id]; !exist {
					reachableVertices[e.Head.Id] = reachableVertices[e.Tail.Id] + 1
					q.QueUp(e.Head)
				}
			}
		}
	}

	return reachableVertices
}

func (x *AdjacencyList) NumVertices() int {
	return len(x.Vertices)
}
func (x *AdjacencyList) NumEdges() int {
	return len(x.Edges)
}
func (x *AdjacencyList) GraphIsDirected() bool {
	return x.IsDirected
}
