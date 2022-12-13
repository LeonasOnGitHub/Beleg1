package graph

import "math"

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
					e := Edge{v2, v1, length}
					p := &e
					x.Edges = append(x.Edges, p)
					if !x.IsDirected {
						e := Edge{v1, v2, length}
						p := &e
						x.Edges = append(x.Edges, p)
					}
				}
			}
		}
	}

}

var reachableVertices = make(map[string]bool)

func (x *AdjacencyList) DFS(nodeId string) map[string]bool {
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
	l := reachableVertices[nodeId]
	for len(q.vertices) > 0 {
		v := q.dequeue()
		for _, e := range x.Edges {
			if e.Tail == v {
				//checks if edging vertex is in the map, if not adds it
				if _, exist := reachableVertices[e.Head.Id]; !exist {
					l = reachableVertices[e.Tail.Id] + 1
					reachableVertices[e.Head.Id] = l
					q.QueUp(e.Head)
				}
			}
		}
	}

	return reachableVertices
}

var topoMap = make(map[string]int)
var curLabel = 0

func (x *AdjacencyList) TopoSort() map[string]int {

	//find a Source Vertex
	for _, v := range x.Vertices {
		if x.isSource(v) {
			curLabel = len(x.DFS(v.Id))
		}
	}

	for _, v := range x.Vertices {
		if _, exist := topoMap[v.Id]; !exist {
			x.dfs_Topo(v)

		}
	}

	return topoMap
}

func (x *AdjacencyList) dfs_Topo(v *Vertex) {
	topoMap[v.Id] = 0
	for _, e := range x.Edges {
		if e.Tail == v {
			if _, exist := topoMap[e.Head.Id]; !exist {
				x.dfs_Topo(e.Head)
			}

		}
	}
	topoMap[v.Id] = curLabel
	curLabel--
}

// return whether the vertex is a valid source vertex for topo
func (x *AdjacencyList) isSource(v *Vertex) bool {
	for _, e := range x.Edges {
		if e.Head.Id == v.Id {
			return false
		}
	}
	// so that no vertices which have no connection aren't added to the map
	for _, e := range x.Edges {
		if e.Tail.Id == v.Id {
			return true
		}
	}
	return false
}
func (x *AdjacencyList) Dijkstra(id string) map[string]float64 {
	dijkMap := make(map[string]float64)
	dijkMap[id] = 0

	for _, e := range x.Edges {
		if _, exist := dijkMap[e.Tail.Id]; exist {
			if _, exist := dijkMap[e.Head.Id]; !exist {
				dijkMap[e.Head.Id] = e.Length + dijkMap[e.Tail.Id]
			} else if e.Length+dijkMap[e.Tail.Id] < dijkMap[e.Head.Id] {
				dijkMap[e.Head.Id] = e.Length + dijkMap[e.Tail.Id]
			}
		}
	}
	return dijkMap
}
func (x AdjacencyList) DijkstaHeap(id string) map[string]float64 {
	dijkMap := make(map[string]float64)
	h := MakeHeap()
	var dis float64
	h.Insert(dis, id)
	for _, v := range x.Vertices {
		if v.Id != id {
			h.Insert(math.Inf(1), v.Id)
		}
	}
	for len(h.tree) > 0 {
		w := *h.ExtractMin()
		dijkMap[w.v] = w.d
		dis = dis + w.d
		for _, edge := range x.Edges {
			if edge.Tail.Id == w.v {
				h.Delete(edge.Head.Id)
				h.Insert(dis+edge.Length, edge.Head.Id)
			}

		}
	}
	return dijkMap
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
