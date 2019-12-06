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

type Search struct {
	seen     *Set
	frontier *Set
}

func find(g *Graph, from, to string, seen, frontier, path *Set) bool {
	fmt.Println("visiting", from)
	if from == to {
		path.Add(to)
		return false
	}

	return true
}

func Find(g *Graph, from, to string) bool {
	path := NewSet()
	find(g, from, to, NewSet(), NewSet(), path)
	fmt.Println(path)
	return false
}

func main() {
	inst.Run("day06p0", loadMap)
	inst.Run("day06p1", part1)
}
