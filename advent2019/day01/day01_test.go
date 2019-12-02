package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	modules := []int{
		12, 14, 1969, 100756,
	}
	expectedFuel := 34241

	fuel := calculateFuel(modules)
	if fuel != expectedFuel {
		t.Fatalf("expected fuel cost %d, have fuel cost %d", expectedFuel, fuel)
	}
}

func TestCalibratedFuelCost(t *testing.T) {
	cases := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for module, expected := range cases {
		fuel := fuelCost(module)
		if fuel != expected {
			t.Fatalf("Module %d should have fuel cost of %d, but have fuel cost %d",
				module, expected, fuel)
		}
	}
}
