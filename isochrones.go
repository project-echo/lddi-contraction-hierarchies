package ch

import (
	"container/heap"
	"fmt"
)

// Isochrones Returns set of vertices and corresponding distances restricted by maximum travel cost for source vertex
// source - source vertex (user defined label)
// maxCost - restriction on travel cost for breadth search
// See ref. https://wiki.openstreetmap.org/wiki/Isochrone and https://en.wikipedia.org/wiki/Isochrone_map
// Note: implemented breadth-first searching path algorithm does not guarantee shortest pathes to reachable vertices (until all edges have cost 1.0). See ref: https://en.wikipedia.org/wiki/Breadth-first_search
// Note: result for estimated costs could be also inconsistent due nature of data structure
func (graph *Graph) Isochrones(source int64, maxCost float64) (map[int64]float64, error) {
	var ok bool
	if source, ok = graph.mapping[source]; !ok {
		return nil, fmt.Errorf("no such source")
	}
	Q := NewMinHeap()
	heap.Init(Q)
	distance := make(map[int64]float64, len(graph.Vertices))
	Q.Push(&MinHeapVertex{ID: source, Distance: 0})
	visit := make(map[int64]bool)
	for Q.Len() != 0 {
		next := heap.Pop(Q).(*MinHeapVertex)
		visit[next.ID] = true
		if next.Distance <= maxCost {
			distance[graph.Vertices[next.ID].Label] = next.Distance
			vertexList := graph.Vertices[next.ID].outIncidentEdges
			for i := range vertexList {
				neighbor := vertexList[i].vertexID
				if v1 := graph.shortcuts[next.ID]; v1 != nil {
					if _, ok2 := v1[neighbor]; ok2 {
						// Ignore shortcut
						continue
					}
				}
				target := vertexList[i].vertexID
				cost := vertexList[i].weight
				alt := distance[graph.Vertices[next.ID].Label] + cost
				if visit[target] {
					continue
				}
				Q.Push(&MinHeapVertex{ID: target, Distance: alt})
			}
		}
	}
	return distance, nil
}
