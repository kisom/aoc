package intcode

import "fmt"

// This is the VM struct I was using at first.

type VM struct {
	mem []int
	ip  int
}

func (vm *VM) Peek(at int) int {
	return vm.mem[at]
}

func (vm *VM) Poke(at int, v int) {
	vm.mem[at] = v
}

func Load(prog []int) *VM {
	vm := &VM{
		mem: make([]int, len(prog)),
	}

	copy(vm.mem, prog)
	return vm
}

func (vm *VM) Run() error {
	vm.ip = 0
	for {
		run, err := vm.step()
		if err != nil {
			return err
		}

		if !run {
			return nil
		}
	}
}

func (vm *VM) step() (bool, error) {
	switch vm.mem[vm.ip] {
	case OpAdd:
		return vm.add()
	case OpMul:
		return vm.mul()
	case OpHalt:
		return false, nil
	}

	return false, fmt.Errorf("intcode: unknown opcode [ip=%d op=%d]",
		vm.ip, vm.mem[vm.ip])
}

func (vm *VM) add() (bool, error) {
	p1 := vm.mem[vm.ip+1]
	p2 := vm.mem[vm.ip+2]
	to := vm.mem[vm.ip+3]

	vm.mem[to] = vm.mem[p1] + vm.mem[p2]
	vm.ip += 4
	return true, nil
}

func (vm *VM) mul() (bool, error) {
	p1 := vm.mem[vm.ip+1]
	p2 := vm.mem[vm.ip+2]
	to := vm.mem[vm.ip+3]

	vm.mem[to] = vm.mem[p1] * vm.mem[p2]
	vm.ip += 4
	return true, nil
}
