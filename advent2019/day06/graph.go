package main

import (
	"strings"
)

// Graph is a directed graph.
type Graph struct {
	nodes   map[string]bool
	edges   map[string]map[string]bool
	reverse map[string]map[string]bool
}

// NewGraph must be used to initialise the graph.
func NewGraph() *Graph {
	return &Graph{
		nodes:   map[string]bool{},
		edges:   map[string]map[string]bool{},
		reverse: map[string]map[string]bool{},
	}
}

// LinksFrom returns all the links from a node as a slice.
func (g *Graph) LinksFrom(node string) (edges []string) {
	for e := range g.edges[node] {
		edges = append(edges, e)
	}
	return
}

// Add registers a node with the graph.
func (g *Graph) Add(node string) {
	g.nodes[node] = true
}

// Link registers an edge with the graph, adding the nodes as needed.
func (g *Graph) Link(obj, com string) {
	if g.edges[obj] == nil {
		g.edges[obj] = map[string]bool{}
	}

	g.edges[obj][com] = true

	if g.reverse[com] == nil {
		g.reverse[com] = map[string]bool{}
	}
	g.reverse[com][obj] = true
}

// Count returns a count of all the edges in the graph.
func (g *Graph) Count() int {
	count := 0
	for node := range g.edges {
		count += g.CountFor(node)
	}
	return count
}

// CountFor returns a count of all of this node's edges.
func (g *Graph) CountFor(node string) int {
	if !g.nodes[node] {
		return 0
	}

	if g.edges[node] == nil {
		return 0
	}

	count := 0
	for node := range g.edges[node] {
		count++
		count += g.CountFor(node)
	}

	return count
}

// Neighbours returns all the neighbours of this string, including any
// parent nodes.
func (g *Graph) Neighbours(node string) []string {
	neighbours := make([]string, 0, len(g.edges[node])+1)
	for k := range g.edges[node] {
		neighbours = append(neighbours, k)
	}

	for direct := range g.reverse[node] {
		neighbours = append(neighbours, direct)
	}

	return neighbours
}

func (g *Graph) AddOrbit(entry string) {
	odata := strings.SplitN(entry, ")", 2)
	com := odata[0]
	obj := odata[1]
	g.Add(com)
	g.Add(obj)
	g.Link(obj, com)
}

func (g *Graph) LoadMap(entries []string) {
	for _, entry := range entries {
		g.AddOrbit(entry)
	}
}

// Search conducts a BFS search over the graph.
func (g *Graph) Search(from, to string) []string {
	frontier := NewSet(from)
	cameFrom := map[string]string{}

	start := from

	for _, neighbour := range g.Neighbours(from) {
		cameFrom[neighbour] = from
		frontier.Add(neighbour)
	}

	frontier.AddList(g.Neighbours(from))

	for !frontier.Empty() {
		node := frontier.Pop()
		if node == to {
			break
		}

		for _, neighbour := range g.Neighbours(node) {
			if !frontier.Has(neighbour) {
				cameFrom[neighbour] = node
				frontier.Add(neighbour)
			}
		}
		from = node
	}

	paths := []string{}
	for to != start {
		paths = append(paths, to)
		to = cameFrom[to]
	}
	paths = append(paths, to)
	for i := 0; i < len(paths)/2; i++ {
		j := len(paths) - i - 1
		paths[i], paths[j] = paths[j], paths[i]
	}

	return paths
}
