/*
Dijkstra is an algorithm for searching the short path between two nodes, visiting the neighbors of each node and calculating the cost and the path from origin node keeping always the smallest value, for that we can use a min-heap to keep the min value in each iteration, using push and pop operation, both operations are O(log n).

1. implement for min-heap, golang has a package in its standard library for that.
2. implement the logic for the graph, for that, we use a struct that contains a map

*/
package main

import (
	hp "container/heap"
	"fmt"
)

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

func main() {
	fmt.Println("Dijkstra")
	// Example
	g := newGraph()
	g.addEdge("S", "B", 4)
	g.addEdge("S", "C", 2)
	g.addEdge("B", "C", 1)
	g.addEdge("B", "D", 5)
	g.addEdge("C", "D", 8)
	g.addEdge("C", "E", 10)
	g.addEdge("D", "E", 2)
	g.addEdge("D", "T", 6)
	g.addEdge("E", "T", 2)
	fmt.Println(g.getPath("S", "T"))
}
