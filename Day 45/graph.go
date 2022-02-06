package main

type Vertex struct {
	// Key is the uid of the vertex
	Key      int
	Vertices map[int]*Vertex
}

// Create a contruction function for Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

type Graph struct {
	//Vertices describes all vertices contained in the graph
	Vertices map[int]*Vertex
	directed bool
}

// Define contructor that create directed or undirected graph

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

//Define a contructor to add vertices
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2, int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}
