package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kisom/aoc/advent2019/inst"
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
	fuelCost := 0

	cost := (module / 3) - 2
	for cost > 0 {
		fuelCost += cost
		cost = cost/3 - 2
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
		fuel += fuelCost(module)
	}

	return fuel
}

func part1() {
	modules, err := readFile("input.txt")
	die.If(err)

	fuel := calculateFuel(modules)
	fmt.Println(fuel)
}

func part2() {
	modules, err := readFile("input.txt")
	die.If(err)

	fuel := calculateFuel2(modules)
	fmt.Println(fuel)
}

func main() {
	inst.Run("day01p1", part1)
	inst.Run("day01p2", part2)
}
