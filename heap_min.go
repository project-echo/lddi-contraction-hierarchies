package ch

import "container/heap"

type minHeapVertex struct {
	id       int64
	distance float64
}

type minHeap struct {
	ids       []int64
	distances map[int64]float64
}

func NewMinHeap() *minHeap {
	return &minHeap{
		ids:       make([]int64, 0),
		distances: make(map[int64]float64),
	}
}

func (h minHeap) Len() int { return len(h.ids) }

func (h minHeap) Less(i, j int) bool {
	a := h.ids[i]
	b := h.ids[j]
	return h.distances[a] < h.distances[b]
}

// Min-Heap
func (h minHeap) Swap(i, j int) { h.ids[i], h.ids[j] = h.ids[j], h.ids[i] }

func (h *minHeap) Push(x interface{}) {
	hv := x.(*minHeapVertex)
	if _, exists := h.distances[hv.id]; !exists {
		h.ids = append(h.ids, hv.id)
	}

	h.distances[hv.id] = hv.distance
}

func (h *minHeap) Pop() interface{} {
	heapSize := len(h.ids)
	lastNode := (h.ids)[heapSize-1]
	h.ids = (h.ids)[0 : heapSize-1]
	delete(h.distances, lastNode)

	return &minHeapVertex{id: lastNode, distance: h.distances[lastNode]}
}

func (h *minHeap) add_with_priority(id int64, val float64) {
	nds := &minHeapVertex{
		id:       id,
		distance: val,
	}
	heap.Push(h, nds)
}
