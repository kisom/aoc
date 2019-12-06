package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kisom/aoc/advent2019/inst"
	"github.com/kisom/goutils/die"
)

var universalOrbitMap []string

const expectedMapSize = 1656

func loadMap() {
	file, err := os.Open("navmap.txt")
	die.If(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		universalOrbitMap = append(universalOrbitMap, strings.TrimSpace(scanner.Text()))
	}

	return
}

func part1() {
	g := NewGraph()
	g.LoadMap(universalOrbitMap)
	if len(universalOrbitMap) != expectedMapSize {
		panic(fmt.Sprintf("universal map should be %d but is %d!",
			expectedMapSize, len(universalOrbitMap)))
	}
	fmt.Println(g.Count())
}

func part2() {
	g := NewGraph()
	g.LoadMap(universalOrbitMap)
	if len(universalOrbitMap) != expectedMapSize {
		panic(fmt.Sprintf("universal map should be %d but is %d!",
			expectedMapSize, len(universalOrbitMap)))
	}

	paths := g.Search(g.LinksFrom("YOU")[0], g.LinksFrom("SAN")[0])
	fmt.Println(len(paths))
}

func main() {
	inst.Run("day06p0", loadMap)
	inst.Run("day06p1", part1)
	inst.Run("day06p2", part2)
}
