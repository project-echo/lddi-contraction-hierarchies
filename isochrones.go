package ch

import (
	"container/heap"
	"fmt"
)

// Isochrones Returns set of vertices and corresponding distances restricted by maximum travel cost for source vertex
// source - source vertex (user defined label)
// maxCost - restriction on travel cost
// See ref. https://wiki.openstreetmap.org/wiki/Isochrone and https://en.wikipedia.org/wiki/Isochrone_map
func (graph *Graph) Isochrones(source int64, maxCost float64) (map[int64]float64, error) {
	var ok bool
	if source, ok = graph.mapping[source]; !ok {
		return nil, fmt.Errorf("no such source")
	}
	Q := &minHeap{}
	heap.Init(Q)
	distance := make(map[int64]float64, len(graph.Vertices))
	distance[graph.Vertices[source].Label] = 0.0
	Q.Push(&minHeapVertex{id: source, distance: 0})
	visit := make(map[int64]bool)
	for Q.Len() != 0 {
		next := heap.Pop(Q).(*minHeapVertex)
		if _, ex := visit[next.id]; ex {
			// ignore nodes already visited with lower distance, as we don't update priority queue when a shorter path is found
			continue
		}
		visit[next.id] = true
		vertexList := graph.Vertices[next.id].outIncidentEdges
		for i := range vertexList {
			neighbor := vertexList[i].vertexID
			if v1 := graph.shortcuts[next.id]; v1 != nil {
				if _, ok2 := v1[neighbor]; ok2 {
					// Ignore shortcut
					continue
				}
			}
			target := vertexList[i].vertexID
			cost := vertexList[i].weight
			alt := distance[graph.Vertices[next.id].Label] + cost
			if alt > maxCost {
				continue
			}
			if visit[target] {
				continue
			}
			if prevDistance, ex := distance[graph.Vertices[target].Label]; !ex || alt < prevDistance {
				distance[graph.Vertices[target].Label] = alt
				Q.Push(&minHeapVertex{id: target, distance: alt})
			}
		}
	}
	return distance, nil
}
