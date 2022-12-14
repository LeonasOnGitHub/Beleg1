package graph

type Key struct {
	d float64 //distance
	v string  //nodeId
}

type Heap struct {
	tree     []*Key
	position map[string]int
}

func MakeHeap() *Heap {
	return &Heap{make([]*Key, 0), make(map[string]int)}

}
func (x *Heap) Insert(dis float64, node string) {
	nextX := Key{dis, node}
	nextK := &nextX
	x.tree = append(x.tree, nextK)
	x.position[nextK.v] = len(x.tree) - 1
	// Swap the new X with his parent until it is in the right spot
	for i := len(x.tree) - 1; i > 0 && nextX.d < x.tree[i/2-1+i%2].d; i = i/2 - 1 + i%2 {
		swap := x.tree[i/2-1+i%2]
		x.tree[i/2-1+i%2] = nextK
		x.tree[i] = swap
		x.position[nextK.v] = i
	}
}

func (x *Heap) ExtractMin() *Key {
	minResult := x.tree[0]

	//swap the last and first element and reduce the array
	x.swap(len(x.tree)-1, 0)
	delete(x.position, x.tree[len(x.tree)-1].v)
	x.tree = x.tree[:len(x.tree)-1]

	x.reorder()

	return minResult
}
func (x *Heap) FindMin() string {
	return x.tree[0].v
}
func (x *Heap) Heapify(arr []*Key) {
	for _, i := range arr {
		x.Insert(i.d, i.v)
	}
}
func (x Heap) Delete(id string) {
	//swap the last and first element and reduce the array
	x.swap(len(x.tree)-1, x.position[id])
	delete(x.position, id)
	x.tree = x.tree[:len(x.tree)-1]

	x.reorder()
}
func (x Heap) reorder() {
	for j := 0; j < len(x.tree)-1; {
		next := true
		//check on the left side to reorder
		if j*2+1 < len(x.tree) && x.tree[j].d > x.tree[j*2+1].d {
			x.swap(j*2+1, j)
			next = false
			//check on the right side to reorder
		} else if j*2+2 < len(x.tree) && x.tree[j].d > x.tree[j*2+2].d {
			x.swap(j*2+2, j)
			next = false
		}
		if next {
			j = j*2 + 1
		}
	}
}
func (x Heap) swap(to, from int) {
	swapValue := x.tree[to]
	x.tree[to] = x.tree[from]
	x.tree[from] = swapValue

	x.position[x.tree[from].v] = from
	x.position[x.tree[to].v] = to
}
