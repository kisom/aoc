package main

import "testing"

func TestClosest(t *testing.T) {
	wire1 := []string{
		"R8", "U5", "L5", "D3",
	}
	wire2 := []string{
		"U7", "R6", "D4", "L4",
	}

	p1 := buildPath(wire1)
	p2 := buildPath(wire2)
	isects := intersections(p1, p2)

	if len(isects) == 0 {
		t.Fatal("no intersections found")
	}

	_, clodst := closest(isects)
	if clodst != 6 {
		t.Fatalf("closest: expected dist=6, have dist=%d", clodst)
	}
}

func TestFastest(t *testing.T) {
	wire1 := []string{
		"R8", "U5", "L5", "D3",
	}
	wire2 := []string{
		"U7", "R6", "D4", "L4",
	}

	p1 := buildPath(wire1)
	p2 := buildPath(wire2)
	isects := intersections(p1, p2)

	if len(isects) == 0 {
		t.Fatal("no intersections found")
	}

	_, cloz := fastest(isects)
	if cloz != 30 {
		t.Fatalf("closest: expected dist=30, have dist=%d", cloz)
	}
}
