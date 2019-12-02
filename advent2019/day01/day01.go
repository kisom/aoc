package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kisom/goutils/die"
)

func readFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	modules := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		modules = append(modules, n)
	}

	return modules, nil
}

func fuelCost(module int) int {
	cost := (module / 3) - 2
	fuelCost := cost
	for cost > 0 {
		cost = cost/3 - 2
		fuelCost += cost
	}

	return fuelCost
}

func calculateFuel(modules []int) int {
	fuel := 0
	for _, module := range modules {
		moduleFuel := (module / 3) - 2
		fuel += moduleFuel
	}

	return fuel
}

func calculateFuel2(modules []int) int {
	fuel := 0
	for _, module := range modules {
		moduleFuel := fuelCost(module)
		fuel += moduleFuel
	}

	return fuel
}

func main() {
	modules, err := readFile("input.txt")
	die.If(err)

	fuel := calculateFuel(modules)
	fmt.Printf("Total fuel: %d\n", fuel)
	fuel = calculateFuel2(modules)
	fmt.Printf("Total calibrated fuel: %d\n", fuel)
}
