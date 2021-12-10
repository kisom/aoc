package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"git.sr.ht/~kisom/goutils/die"
)

func part1(r io.Reader) (int, error) {
	prevDepth := -1
	increases := 0

	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()
		depth, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}

		if prevDepth >= 0 {
			if depth > prevDepth {
				increases++
			}
		}
		prevDepth = depth
	}

	return increases, nil
}

func sum(ns []int) int {
	if len(ns) != 3 {
		panic(fmt.Sprintf("length %d isn't 3", len(ns)))
	}
	sum := 0
	for i := 0; i < len(ns); i++ {
		sum += ns[i]
	}
	return sum
}

func part2(r io.Reader) (int, error) {
	depths := []int{}
	s := bufio.NewScanner(r)
	count := 0
	
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return -1, err
		}
		depths = append(depths, n)
	}

	for i := 3; i < len(depths); i++ {
		if sum(depths[i-3:i]) < sum(depths[i-2:i+1]) {
			count++
		}
	}

	return count, nil
}

func main() {
	file, err := os.Open("input1.txt")
	die.If(err)
	defer file.Close()

	di, err := part1(file)
	die.If(err)

	fmt.Println("part a:", di)

	file.Seek(0, 0)

	di, err = part2(file)
	fmt.Println("part b:", di)
}
