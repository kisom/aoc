package ic

import (
	"bytes"
	"strings"
	"testing"
)

func cmpis(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type vmTest struct {
	prog  []int
	final []int
	mods  []mod
}

var vmTestCases = []vmTest{
	{
		prog: []int{
			1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50,
		},
		final: []int{
			3500, 9, 10, 70,
			2, 3, 11, 0,
			99,
			30, 40, 50,
		},
	},
	{
		prog:  []int{1, 0, 0, 0, 99},
		final: []int{2, 0, 0, 0, 99},
	},
	{
		prog:  []int{2, 3, 0, 3, 99},
		final: []int{2, 3, 0, 6, 99},
	},
	{
		prog:  []int{2, 4, 4, 5, 99, 0},
		final: []int{2, 4, 4, 5, 99, 9801},
	},
	{
		prog:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		final: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
	{
		prog:  []int{0, 0, 1, -1, 99, 5, 6, 0, 99},
		final: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		mods: []mod{
			Mod(0, 1), Mod(1, 1), Mod(3, 4),
		},
	},
}

func TestVM(t *testing.T) {
	for _, c := range vmTestCases {
		prog, err := Run(c.prog, nil, c.mods...)
		if err != nil {
			t.Error(err)
		}

		if !cmpis(prog, c.final) {
			t.Errorf("VM is not in an expected state:\nexp: %+v\ncur: %+v",
				c.final, prog)
		}
	}
}

func TestErr(t *testing.T) {
	prog := []int{-1, 1, 1, 1}
	_, err := Run(prog, nil)
	if err == nil {
		t.Error("expected an error with opcode=-1")
	}
}

func TestIO(t *testing.T) {
	expectations := map[string]string{
		"1":  "? 2",
		"-1": "? 0",
	}
	// This program inputs a number n and outputs n+1.
	prog := []int{3, 4, 1101, 1, 0, 5, 4, 5, 99}
	for input, expectation := range expectations {
		buf := bytes.NewBufferString(input + "\n")
		_, err := Run(prog, buf)
		if err != nil {
			t.Error(err)
		}

		if output := strings.TrimSpace(buf.String()); output != expectation {
			t.Errorf("Run: expected output '%s', have '%s'", expectation, output)
		}
	}
}

// From day05.txt:
// Here's a larger example:
//
//   3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
//   1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
//   999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99
//
// The above example program uses an input instruction to ask for a
// single number. The program will then output 999 if the input value is
// below 8, output 1000 if the input value is equal to 8, or output 1001
// if the input value is greater than 8.
func TestIs8(t *testing.T) {
	expectations := map[string]string{
		"1":  "? 999",
		"7":  "? 999",
		"8":  "? 1000",
		"9":  "? 1001",
		"11": "? 1001",
	}
	prog := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20,
		1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125,
		20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101,
		1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}

	for input, expectation := range expectations {
		buf := bytes.NewBufferString(input + "\n")
		_, err := Run(prog, buf)
		if err != nil {
			t.Fatal(err)
		}

		if output := strings.TrimSpace(buf.String()); output != expectation {
			t.Fatalf("Run: expected output '%s', have '%s'", expectation, output)
		}
	}
}
