package tests

import (
	"Beleg1/graph"
	"testing"
)

func TestHeapAlreadyRightOrder(t *testing.T) {
	h := graph.Heap{}
	arr := []int{1, 2, 3, 4, 5, 6}
	for _, i := range arr {
		h.Insert(i)
	}
	if h.ExtractMin() != 1 {
		t.Error("It should be 1 but it is ", h.ExtractMin())
	}
}
func TestHeaSmall(t *testing.T) {
	h := graph.Heap{}
	arr := []int{3, 1, 3}
	for _, i := range arr {
		h.Insert(i)
	}
	if h.ExtractMin() != 1 {
		t.Error("It should be 1 but it is ", h.ExtractMin())
	}
}
func TestHeapBig(t *testing.T) {
	h := graph.Heap{}
	arr := []int{3, 1, 3, 0, 0, 1, 5, 6, 7, 8, 4, 7, 8, 5, 7, 8, 5, 11, 55}
	for _, i := range arr {
		h.Insert(i)
	}
	if h.ExtractMin() != 0 {
		t.Error("It should be 0 but it is ", h.ExtractMin())
	}
}
func TestHeapWithNegative(t *testing.T) {
	h := graph.Heap{}
	arr := []int{3, 1, -3, 0, 5, -1}
	for _, i := range arr {
		h.Insert(i)
	}
	if h.ExtractMin() != -3 {
		t.Error("It should be -3 but it is ", h.ExtractMin())
	}
}
