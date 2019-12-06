package main

import (
	"bytes"
	"encoding/json"
	"strings"
)

type Graph struct {
	Nodes  map[string]bool
	Edges  map[string]map[string]bool
	Direct map[string]string
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:  map[string]bool{},
		Edges:  map[string]map[string]bool{},
		Direct: map[string]string{},
	}
}

func (g *Graph) String() string {
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	json.Indent(buf, data, "", "\t")
	return buf.String()
}

func (g *Graph) Add(node string) {
	g.Nodes[node] = true
}

func (g *Graph) Link(obj, com string) {
	if g.Edges[obj] == nil {
		g.Edges[obj] = map[string]bool{}
	}

	g.Edges[obj][com] = true
	g.Direct[com] = obj
}

func (g *Graph) Count() int {
	count := 0
	for node := range g.Edges {
		count += g.CountFor(node)
	}
	return count
}

func (g *Graph) CountFor(node string) int {
	if !g.Nodes[node] {
		return 0
	}

	if g.Edges[node] == nil {
		return 0
	}

	count := 0
	for node := range g.Edges[node] {
		count++
		count += g.CountFor(node)
	}

	return count
}

func (g *Graph) Neighbours(node string) []string {
	neighbours := make([]string, 0, len(g.Edges[node])+1)
	for k := range g.Edges[node] {
		neighbours = append(neighbours, k)
	}

	if direct := g.Direct[node]; direct != "" {
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
