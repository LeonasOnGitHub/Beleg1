package graph

type Queue struct {
	vertices []*Vertex
}

func (x *Queue) QueUp(nextQ *Vertex) {
	x.vertices = append(x.vertices, nextQ)
}

func (x *Queue) dequeue() *Vertex {
	v := x.vertices[0]
	var s []*Vertex = x.vertices[1:len(x.vertices)]
	x.vertices = s
	return v
}
