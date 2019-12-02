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

func isVMinExpectedState(vm *VM, final []int) bool {
	return cmpis(vm.mem, final)
}

func runTestVM(prog []int, final []int, t *testing.T) {
	vm := Load(prog)
	err := vm.Run()
	if err != nil {
		t.Error(err)
	}

	if !isVMinExpectedState(vm, final) {
		t.Errorf("VM is not in an expected state:\nexp: %+v\ncur: %+v",
			final, vm.mem)
	}
}

type vmTest struct {
	prog  []int
	final []int
}

var vmTestCases = []vmTest{
	vmTest{
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
	vmTest{
		prog:  []int{1, 0, 0, 0, 99},
		final: []int{2, 0, 0, 0, 99},
	},
	vmTest{
		prog:  []int{2, 3, 0, 3, 99},
		final: []int{2, 3, 0, 6, 99},
	},
	vmTest{
		prog:  []int{2, 4, 4, 5, 99, 0},
		final: []int{2, 4, 4, 5, 99, 9801},
	},
	vmTest{
		prog:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		final: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
}

func TestVM(t *testing.T) {
	for _, c := range vmTestCases {
		prog, err := Run(c.prog)
		if err != nil {
			t.Error(err)
		}

		if !cmpis(prog, c.final) {
			t.Errorf("VM is not in an expected state:\nexp: %+v\ncur: %+v",
				c.final, prog)
		}
		runTestVM(c.prog, c.final, t)
	}
}
