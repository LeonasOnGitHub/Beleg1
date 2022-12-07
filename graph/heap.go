package graph

type Heap struct {
	tree []int
}

func (x *Heap) Insert(nextX int) {
	x.tree = append(x.tree, nextX)
	// Swap the new X with his parent until it is in the right spot
	for i := len(x.tree) - 1; i > 0 && nextX < x.tree[i/2-1+i%2]; i = i/2 - 1 + i%2 {
		swap := x.tree[i/2-1+i%2]
		x.tree[i/2-1+i%2] = nextX
		x.tree[i] = swap
	}
}

func (x *Heap) ExtractMin() int {
	minResult := x.tree[0]

	//swap the last and first element and reduce the array
	swap := x.tree[len(x.tree)-1]
	x.tree[0] = swap
	x.tree = x.tree[:len(x.tree)-2]

	//reorder the array
	for j := 0; j < len(x.tree)-1; j = j*2 + 1 {
		//check on the left side to reorder
		if j*2+1 < len(x.tree) && x.tree[j] > x.tree[j*2+1] {
			swap = x.tree[j*2+1]
			x.tree[j*2+1] = x.tree[j]
			x.tree[j] = swap
			//check on the right side to reorder
		} else if j*2+2 < len(x.tree) && x.tree[j] > x.tree[j*2+2] {
			swap = x.tree[j*2+2]
			x.tree[j*2+2] = x.tree[j]
			x.tree[j] = swap
		} else {
			break // break the loop if the order is correct
		}
	}
	return minResult
}
