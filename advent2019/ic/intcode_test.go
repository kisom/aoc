package ic

import "testing"

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
		prog, err := Run(c.prog, c.mods...)
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
	_, err := Run(prog)
	if err == nil {
		t.Error("expected an error with opcode=-1")
	}
}
