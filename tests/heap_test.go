package tests

import (
	"fmt"
	"testing"
)

func TestDijkstraHeap(t *testing.T) {
	resultDijk := alD.DijkstaHeap("A")
	var dummyMapDist = make(map[string]float64)
	dummyMapDist["A"] = 0
	dummyMapDist["B"] = 0.5
	dummyMapDist["C"] = 0.9
	dummyMapDist["D"] = 1.2

	for key, value := range resultDijk {
		if value != dummyMapDist[key] {
			t.Error("The Vertex ", key, " should have the length", dummyMapDist[key], " but has ", value)
		}
	}
}
func TestProblem98Small(t *testing.T) {
	testGraph := initGraph9("smallTest.txt", t)
	resultDijk := testGraph.DijkstaHeap("1")
	for s, f := range resultDijk {
		fmt.Println(s, ": ", f)

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////			 	OLD HEAP TESTS					////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*func TestHeapAlreadyRightOrder(t *testing.T) {
	h := new(graph.Heap)
	arrD := []float64{1, 2, 3, 4, 5, 6}
	arrV := []string{"1", "2", "3", "4", "5", "6"}
	for i2, i := range arrD {
		h.Insert(i, arrV[i2])
	}
	if h.ExtractMin() != "1" {
		t.Error("It should be 1 but it is ", h.ExtractMin())
	}
}
func TestHeaSmall(t *testing.T) {
	h := graph.Heap{}
	arrD := []float64{3, 1, 3}
	arrV := []string{"1", "2", "3"}
	for i2, i := range arrD {
		h.Insert(i, arrV[i2])
	}
	if h.ExtractMin() != "2" {
		t.Error("It should be 2	 but it is ", h.ExtractMin())
	}
}
func TestHeapBig(t *testing.T) {
	h := graph.Heap{}
	arrD := []float64{3, 1, 3, 0, 0, 1, 5, 6, 7, 8, 4, 7, 8, 5, 7, 8, 5, 11, 55}
	arrV := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
	for i2, i := range arrD {
		h.Insert(i, arrV[i2])
	}
	if h.ExtractMin() != "4" || h.ExtractMin() != "5" {
		t.Error("It should be 4 or 5 but it is ", h.ExtractMin())
	}
	if h.ExtractMin() != "4" || h.ExtractMin() != "5" {
		t.Error("It should be 4 or 5 but it is ", h.ExtractMin())
	}
}
func TestHeapWithNegative(t *testing.T) {
	h := graph.Heap{}
	arrD := []float64{3, 1, -3, 0, 5, -1}
	arrV := []string{"1", "2", "3", "4", "5", "6"}
	for i2, i := range arrD {
		h.Insert(i, arrV[i2])
	}
	if h.ExtractMin() != "3" {
		t.Error("It should be 3 but it is ", h.ExtractMin())
	}
	if h.ExtractMin() != "6" {
		t.Error("It should be 6 but it is ", h.ExtractMin())
	}

}
*/
