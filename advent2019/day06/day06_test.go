package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	mapData := []string{"COM)B", "B)C", "C)D"}
	g := NewGraph()
	g.LoadMap(mapData)

	fmt.Println(g)
	if count := g.CountFor("D"); count != 3 {
		t.Errorf("expect 3 orbits for D, have %d", count)
	}

	if !g.Edges["B"]["COM"] {
		t.Fatal("expect B to orbit COM")
	}

	if !g.Edges["C"]["B"] {
		t.Fatal("expect C to orbit B")
	}

	if !g.Edges["D"]["C"] {
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

	t.Log(g)
}

func TestLoadNavMap(t *testing.T) {
	loadMap()
	if len(universalOrbitMap) != 1656 {
		t.Fatalf("expected 1656 entries in the universal map, have %d",
			len(universalOrbitMap))
	}
}

func TestNeighbours(t *testing.T) {
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

	neighbours := g.Neighbours("K")
	if len(neighbours) != 2 {
		t.Errorf("expected K to have 2 neighbours, have %d neighbours", len(neighbours))
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
		"K)YOU",
		"I)SAN",
	}
	g := NewGraph()
	g.LoadMap(mapData)

	Find(g, "YOU", "SAN")
}
