package main

import (
	"math/rand"
	"testing"
)

func newTestMap() *Graph {
	mapData := []string{
		"K)YOU",
		"I)SAN",
		"K)L",
		"J)K",
		"E)J",
		"D)I",
		"G)H",
		"B)G",
		"E)F",
		"D)E",
		"C)D",
		"B)C",
		"COM)B",
	}

	rand.Shuffle(len(mapData), func(i, j int) {
		mapData[j], mapData[i] = mapData[i], mapData[j]
	})

	g := NewGraph()
	g.LoadMap(mapData)
	return g
}

func TestBasic(t *testing.T) {
	mapData := []string{"COM)B", "B)C", "C)D"}
	g := NewGraph()
	g.LoadMap(mapData)

	if count := g.CountFor("D"); count != 3 {
		t.Errorf("expect 3 orbits for D, have %d", count)
	}

	if !g.edges["B"]["COM"] {
		t.Fatal("expect B to orbit COM")
	}

	if !g.edges["C"]["B"] {
		t.Fatal("expect C to orbit B")
	}

	if !g.edges["D"]["C"] {
		t.Fatal("expect D to orbit C")
	}
}

func TestFullPart1Reversed(t *testing.T) {
	mapData := []string{
		"K)L",
		"J)K",
		"E)J",
		"D)I",
		"G)H",
		"B)G",
		"E)F",
		"D)E",
		"C)D",
		"B)C",
		"COM)B",
	}

	g := NewGraph()
	g.LoadMap(mapData)

	count := g.Count()
	if count != 42 {
		t.Fatalf("expect 42 orbits, have %d\n%#v", count, *g)
	}

	if count = g.CountFor("D"); count != 3 {
		t.Fatalf("expect 3 orbits for D, have %d", count)
	}

	if count = g.CountFor("L"); count != 7 {
		t.Fatalf("expect 7 orbits for D, have %d", count)
	}

	if count = g.CountFor("COM"); count != 0 {
		t.Fatalf("expect 0 orbits for COM, have %d", count)
	}
}

func TestFullPart1(t *testing.T) {
	mapData := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	g := NewGraph()
	g.LoadMap(mapData)

	count := g.Count()
	if count != 42 {
		t.Fatalf("expect 42 orbits, have %d\n%#v", count, *g)
	}

	if count = g.CountFor("D"); count != 3 {
		t.Fatalf("expect 3 orbits for D, have %d", count)
	}

	if count = g.CountFor("L"); count != 7 {
		t.Fatalf("expect 7 orbits for D, have %d", count)
	}

	if count = g.CountFor("COM"); count != 0 {
		t.Fatalf("expect 0 orbits for COM, have %d", count)
	}
}

func TestLoadNavMap(t *testing.T) {
	loadMap()
	if len(universalOrbitMap) != 1656 {
		t.Fatalf("expected 1656 entries in the universal map, have %d",
			len(universalOrbitMap))
	}
}

func TestNeighbours(t *testing.T) {
	g := newTestMap()
	neighbours := g.Neighbours("K")
	if len(neighbours) != 3 {
		t.Errorf("expected K to have 2 neighbours, have %d neighbours [%v]",
			len(neighbours), neighbours)
	}

	m := map[string]bool{}
	for i := range neighbours {
		m[neighbours[i]] = true
	}

	if !m["J"] {
		t.Error("J should be a neighbour of K")
	}

	if !m["L"] {
		t.Errorf("L should be a neighbour of K")
	}
}

func TestFindPath(t *testing.T) {
	for i := 0; i < 8; i++ {
		g := newTestMap()
		expected := []string{
			"YOU", "K", "J", "E", "D", "I", "SAN",
		}

		paths := g.Search("YOU", "SAN")
		if len(paths) != len(expected) {
			t.Fatalf("expected %d nodes in path, have %d [%v]",
				len(expected), len(paths), paths)
		}
		for i := range paths {
			if paths[i] != expected[i] {
				t.Errorf("expected node %d to be %s, but have %s",
					i, expected[i], paths[i])
			}
		}
	}
}
