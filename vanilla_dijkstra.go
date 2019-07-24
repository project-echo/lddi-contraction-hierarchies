package ch

import (
	"container/heap"
	"log"
	"math"
)

// VanillaShortestPath Computes and returns shortest path and it's cost (vanilla Dijkstra's algorithm)
//
// If there are some errors then function returns '-1.0' as cost and nil as shortest path
//
// source User's definied ID of source vertex
// target User's definied ID of target vertex
//
func (graph *Graph) VanillaShortestPath(source, target int) (float64, []int) {

	if source == target {
		return 0, []int{source}
	}
	ok := true

	if source, ok = graph.mapping[source]; !ok {
		log.Println("No such source")
		return -1.0, nil
	}
	if target, ok = graph.mapping[target]; !ok {
		log.Println("No such target")
		return -1.0, nil
	}

	prev := make(map[int]int)

	queryDist := make([]float64, len(graph.Vertices), len(graph.Vertices))
	revDistance := make([]float64, len(graph.Vertices), len(graph.Vertices))

	for i := range queryDist {
		queryDist[i] = math.MaxFloat64
		revDistance[i] = math.MaxFloat64
	}
	queryDist[source] = 0
	revDistance[target] = 0

	forwQ := &forwardPropagationHeap{}
	heap.Init(forwQ)

	heapSource := simpleNode{
		id:          source,
		queryDist:   0,
		revDistance: math.MaxFloat64,
	}
	heap.Push(forwQ, heapSource)

	estimate := math.MaxFloat64

	for forwQ.Len() != 0 {
		if forwQ.Len() != 0 {
			vertex1 := heap.Pop(forwQ).(simpleNode)
			if vertex1.queryDist <= estimate {
				graph.vanillaRelaxEdge(&vertex1, forwQ, prev, queryDist, revDistance)
			}
			if vertex1.queryDist < estimate {
				estimate = vertex1.queryDist + vertex1.revDistance
			}
		}
	}

	if estimate == math.MaxFloat64 {
		return -1.0, nil
	}

	// Compute path
	var path []int
	u := target
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		temp := make([]int, len(path)+1)
		temp[0] = u
		copy(temp[1:], path)
		path = temp
		u = prev[u]
	}
	temp := make([]int, len(path)+1)
	temp[0] = source
	copy(temp[1:], path)
	path = temp

	return estimate, path
}

// vanillaRelaxEdge Edge relaxation
func (graph *Graph) vanillaRelaxEdge(vertex *simpleNode, forwQ *forwardPropagationHeap, prev map[int]int, distances []float64, prevReverse []float64) {
	vertexList := graph.Vertices[vertex.id].outEdges
	costList := graph.Vertices[vertex.id].outECost
	for i := 0; i < len(vertexList); i++ {
		temp := vertexList[i]
		cost := costList[i]
		alt := distances[vertex.id] + cost
		if distances[temp] > alt {
			distances[temp] = alt
			prev[temp] = vertex.id
			node := simpleNode{
				id:          temp,
				queryDist:   alt,
				revDistance: prevReverse[temp],
			}
			heap.Push(forwQ, node)
		}
	}
}