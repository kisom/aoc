package main

import "testing"

/*
   111111 meets these criteria (double 11, never decreases).
   223450 does not meet these criteria (decreasing pair of digits 50).
   123789 does not meet these criteria (no double).
*/
var testNums1 = []int{111111, 223450, 123789}

func TestTwoAdjacent(t *testing.T) {
	if !twoAdjacent(testNums1[0]) {
		t.Errorf("%d should have adjacency", testNums1[0])
	}

	if !twoAdjacent(testNums1[1]) {
		t.Errorf("%d should have adjacency", testNums1[1])
	}

	if twoAdjacent(testNums1[2]) {
		t.Errorf("%d shouldn't have adjacency", testNums1[2])
	}
}

var testNums2 = []int{112233, 123444, 111122}

func TestOnlyTwoAdjacent(t *testing.T) {
	if !onlyTwoAdjacent(testNums2[0]) {
		t.Errorf("%d should have strict pair adjacency", testNums2[0])
	}

	if onlyTwoAdjacent(testNums2[1]) {
		t.Errorf("%d shouldn't have strict pair adjacency", testNums2[1])
	}

	if !onlyTwoAdjacent(testNums2[2]) {
		t.Errorf("%d should have strict pair adjacency", testNums2[2])
	}
}

func TestIsMonotic(t *testing.T) {
	if !isMonotic(testNums1[0]) {
		t.Errorf("%d should be monotonic", testNums1[0])
	}

	if isMonotic(testNums1[1]) {
		t.Errorf("%d shouldn't be monotonic", testNums1[1])
	}

	if !isMonotic(testNums1[2]) {
		t.Errorf("%d should be monotonic", testNums1[2])
	}
}
