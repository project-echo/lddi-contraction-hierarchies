package ch

import "container/heap"

type MinHeapVertex struct {
	ID       int64
	Distance float64
}

type MinHeap struct {
	ids       []int64
	distances map[int64]float64
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		ids:       make([]int64, 0),
		distances: make(map[int64]float64),
	}
}

func (h MinHeap) Len() int { return len(h.ids) }

func (h MinHeap) Less(i, j int) bool {
	a := h.ids[i]
	b := h.ids[j]
	return h.distances[a] < h.distances[b]
}

// Min-Heap
func (h MinHeap) Swap(i, j int) { h.ids[i], h.ids[j] = h.ids[j], h.ids[i] }

func (h *MinHeap) Push(x interface{}) {
	hv := x.(*MinHeapVertex)
	if _, exists := h.distances[hv.ID]; !exists {
		h.ids = append(h.ids, hv.ID)
	}

	h.distances[hv.ID] = hv.Distance
}

func (h *MinHeap) Pop() interface{} {
	heapSize := len(h.ids)
	lastNode := h.ids[heapSize-1]
	lastDistance := h.distances[lastNode] // Capture the distance before deletion
	h.ids = h.ids[:heapSize-1]
	delete(h.distances, lastNode)

	return &MinHeapVertex{ID: lastNode, Distance: lastDistance} // Use the captured distance
}

func (h *MinHeap) add_with_priority(id int64, val float64) {
	nds := &MinHeapVertex{
		ID:       id,
		Distance: val,
	}
	heap.Push(h, nds)
}
