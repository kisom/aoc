package main

import (
	"bytes"
	"fmt"

	"github.com/kisom/aoc/advent2019/inst"
)

var minRange = [6]byte{2, 3, 5, 7, 4, 0} // dec by 1
var maxRange = [6]byte{7, 0, 6, 9, 4, 9} // inc by 1

type counter struct {
	v   [6]byte
	max [6]byte
}

func (c *counter) next() bool {
	for i := 5; i >= 0; i-- {
		c.v[i]++
		if c.v[i] != 0 {
			break
		}
	}

	if bytes.Equal(c.v[:], c.max[:]) {
		return false
	}

	return true
}

func newCounter(min, max [6]byte) *counter {
	c := &counter{}
	copy(c.v[:], min[:])
	copy(c.max[:], max[:])
	return c
}

func twoAdjacent(c *counter) bool {
	for i := 1; i < 6; i++ {
		if c.v[i-1] == c.v[i] {
			return true
		}
	}

	return false
}

func onlyTwoAdjacent(c *counter) bool {
	run := c.v[0]
	count := 1
	ok := false

	for i := 1; i < 6; i++ {
		if c.v[i] == run {
			count++
		} else {
			if count == 2 {
				ok = true
			}
			run = c.v[i]
			count = 1
		}
	}

	if count == 2 {
		ok = true
	}
	return ok
}

func isMonotic(c *counter) bool {
	// relies on the fact that ASCII numerics are digit + 0x30, so
	// they can be compared directly without parsing.
	for i := 1; i < 6; i++ {
		if c.v[i] < c.v[i-1] {
			return false
		}
	}

	return true
}

func findCandidate1(min, max [6]byte) int {
	count := 0
	counter := newCounter(min, max)
	for counter.next() {
		if twoAdjacent(counter) && isMonotic(counter) {
			count++
		}
	}

	return count
}

func findCandidate2(min, max [6]byte) int {
	count := 0
	counter := newCounter(min, max)
	for counter.next() {
		if onlyTwoAdjacent(counter) && isMonotic(counter) {
			count++
		}
	}

	return count
}

func part1() {
	fmt.Println(findCandidate1(minRange, maxRange))
}

func part2() {
	fmt.Println(findCandidate2(minRange, maxRange))
}

func main() {
	inst.Run("day04p1", part1)
	inst.Run("day04p2", part2)
}
