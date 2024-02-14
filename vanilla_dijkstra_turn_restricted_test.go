package ch

import "testing"

type V struct {
	from   int64
	to     int64
	weight float64
}

func TestVanillaTurnRestrictedShortestPathUnidirectional(t *testing.T) {
	vertices := []V{
		{from: 1, to: 2, weight: 1.0},
		{from: 2, to: 3, weight: 3.0},
		{from: 3, to: 4, weight: 1.0},
		{from: 4, to: 5, weight: 1.0},
		{from: 5, to: 6, weight: 1.0},
		{from: 5, to: 7, weight: 1.0},
		{from: 2, to: 5, weight: 1.0},
		{from: 8, to: 2, weight: 1.0},
	}

	graph := Graph{}
	for i := range vertices {
		err := graph.CreateVertex(vertices[i].from)
		if err != nil {
			t.Error(err)
			return
		}
		err = graph.CreateVertex(vertices[i].to)
		if err != nil {
			t.Error(err)
			return
		}
		err = graph.AddEdge(vertices[i].from, vertices[i].to, vertices[i].weight)
		if err != nil {
			t.Error(err)
			return
		}
	}

	graph.PrepareContractionHierarchies()

	restrictions := make(map[int64]map[int64]bool)

	// empty restrictions must calculate the default shortest path
	ans, path := graph.VanillaTurnRestrictedShortestPath(1, 5, restrictions)
	rightPath := []int64{1, 2, 5}
	if len(path) != 3 {
		t.Errorf("Run 1: num of vertices in path should be 5, but got %d %v", len(path), path)
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 1: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}
	if ans != 2 {
		t.Errorf("Run 1: length of path should be 2, but got %f", ans)
		return
	}

	restrictions = make(map[int64]map[int64]bool)
	restrictions[2] = make(map[int64]bool)
	restrictions[2][3] = true
	restrictions[1] = make(map[int64]bool)
	restrictions[1][2] = true
	// closing the middle of the road must return a empty path (empty path only consists from source)
	_, path = graph.VanillaTurnRestrictedShortestPath(1, 5, restrictions)
	rightPath = []int64{1}
	if len(path) != 1 {
		t.Errorf("Run 1: num of vertices in path should be 5, but got %d %v", len(path), path)
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 1: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}

	restrictions = make(map[int64]map[int64]bool)
	restrictions[2] = make(map[int64]bool)
	restrictions[2][5] = true
	ans, path = graph.VanillaTurnRestrictedShortestPath(2, 7, restrictions)
	rightPath = []int64{2, 3, 4, 5, 7}
	if len(path) != 5 {
		t.Errorf("Run 2: num of vertices in path should be 5, but got %d %v", len(path), path)
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 2: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}
	if ans != 6 {
		t.Errorf("Run 2: length of path should be 6, but got %f", ans)
		return
	}

	ans, path = graph.VanillaTurnRestrictedShortestPath(1, 7, restrictions)
	rightPath = []int64{1, 2, 3, 4, 5, 7}
	if len(path) != 6 {
		t.Errorf("Run 3: num of vertices in path should be 6, but got %d", len(path))
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 3: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}
	if ans != 7 {
		t.Errorf("Run 3: length of path should be 7, but got %f", ans)
		return
	}

	t.Log("TestVanillaTurnRestrictedShortestPath is Ok!")
}

func TestVanillaTurnRestrictedShortestPathBidirectional(t *testing.T) {
	// vertices represents a circular graph where each vertex is connected bidirectionally to its adjacent vertices,
	// allowing bidirectional travel on the graph.
	vertices := []V{
		{from: 1, to: 2, weight: 1.0}, {from: 2, to: 1, weight: 1.0},
		{from: 2, to: 3, weight: 1.0}, {from: 3, to: 2, weight: 1.0},
		{from: 3, to: 4, weight: 1.0}, {from: 4, to: 3, weight: 1.0},
		{from: 4, to: 5, weight: 1.0}, {from: 5, to: 4, weight: 1.0},
		{from: 5, to: 6, weight: 1.0}, {from: 6, to: 5, weight: 1.0},
		{from: 6, to: 7, weight: 1.0}, {from: 7, to: 6, weight: 1.0},
		{from: 7, to: 8, weight: 1.0}, {from: 8, to: 7, weight: 1.0},
		{from: 8, to: 9, weight: 1.0}, {from: 9, to: 8, weight: 1.0},
		{from: 9, to: 10, weight: 1.0}, {from: 10, to: 9, weight: 1.0},
		{from: 10, to: 11, weight: 1.0}, {from: 11, to: 10, weight: 1.0},
		{from: 11, to: 12, weight: 1.0}, {from: 12, to: 11, weight: 1.0},
		{from: 12, to: 13, weight: 1.0}, {from: 13, to: 12, weight: 1.0},
		{from: 13, to: 14, weight: 1.0}, {from: 14, to: 13, weight: 1.0},
		{from: 14, to: 15, weight: 1.0}, {from: 15, to: 14, weight: 1.0},
		{from: 15, to: 16, weight: 1.0}, {from: 16, to: 15, weight: 1.0},
		{from: 16, to: 17, weight: 1.0}, {from: 17, to: 16, weight: 1.0},
		{from: 17, to: 18, weight: 1.0}, {from: 18, to: 17, weight: 1.0},
		{from: 18, to: 19, weight: 1.0}, {from: 19, to: 18, weight: 1.0},
		{from: 19, to: 20, weight: 1.0}, {from: 20, to: 19, weight: 1.0},
		{from: 20, to: 1, weight: 1.0}, {from: 1, to: 20, weight: 1.0},
	}

	graph := Graph{}
	for i := range vertices {
		err := graph.CreateVertex(vertices[i].from)
		if err != nil {
			t.Error(err)
			return
		}
		err = graph.CreateVertex(vertices[i].to)
		if err != nil {
			t.Error(err)
			return
		}
		err = graph.AddEdge(vertices[i].from, vertices[i].to, vertices[i].weight)
		if err != nil {
			t.Error(err)
			return
		}
	}

	graph.PrepareContractionHierarchies()

	_, path := graph.VanillaTurnRestrictedShortestPath(10, 1, nil)
	// by default, this path should be calculated shortest path which is 9 steps
	rightPath := []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if len(path) != 10 {
		t.Errorf("Run 1: num of vertices in path should be 10, but got %d %v", len(path), path)
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 1: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}

	// now lets close the shortest path and force algorithm to take a longer turn
	restrictions := make(map[int64]map[int64]bool)
	restrictions[5] = make(map[int64]bool)
	restrictions[5][6] = true

	restrictions[6] = make(map[int64]bool)
	restrictions[6][5] = true

	_, path = graph.VanillaTurnRestrictedShortestPath(10, 1, restrictions)
	//now it should take the longer turn which is 11 steps since we closed the shorter path
	rightPath = []int64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1}
	if len(path) != 12 {
		t.Errorf("Run 2: num of vertices in path should be 12, but got %d %v", len(path), path)
		return
	}
	for i := range path {
		if path[i] != rightPath[i] {
			t.Errorf("Run 2: vertex in path should be %d, but got %d", path[i], rightPath[i])
			return
		}
	}

	// now lets close the both shorter and longer paths and force algorithm to fail
	restrictions = make(map[int64]map[int64]bool)
	restrictions[5] = make(map[int64]bool)
	restrictions[5][6] = true

	restrictions[6] = make(map[int64]bool)
	restrictions[6][5] = true

	restrictions[15] = make(map[int64]bool)
	restrictions[15][16] = true

	restrictions[16] = make(map[int64]bool)
	restrictions[16][15] = true

	cost, path := graph.VanillaTurnRestrictedShortestPath(10, 1, restrictions)
	if len(path) != 1 {
		t.Errorf("Run 3: num of vertices in path should be 1, but got %d %v", len(path), path)
		return
	}

	if cost != Infinity {
		t.Errorf("Unreachable paths cost must be infinite, but VanillaTurnRestrictedShortestPath returned %f cost for unreachable path", cost)
	}
}
