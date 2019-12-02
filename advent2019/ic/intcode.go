package ic

import "fmt"

const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
)

type mod struct {
	Pos int
	Val int
}

// A mod encodes a change to the program to be made at runtime.
func Mod(pos, val int) mod {
	return mod{
		Pos: pos,
		Val: val,
	}
}

// Run copies the program to a new memory space, applies any mods, and
// the runs it, returning the program. For Day 2, all of the results
// are in memory position 0, but I don't want to make that assumption
// for the future.
func Run(prog []int, mods ...mod) ([]int, error) {
	mem := make([]int, len(prog))
	copy(mem, prog)

	for _, mod := range mods {
		mem[mod.Pos] = mod.Val
	}

	ip := 0
	run := true
	for {
		if !run {
			break
		}

		op := mem[ip]
		switch op {
		case OpAdd:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			dest := mem[ip+3]
			mem[dest] = mem[arg1] + mem[arg2]
			ip += 4
		case OpMul:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			dest := mem[ip+3]
			mem[dest] = mem[arg1] * mem[arg2]
			ip += 4
		case OpHalt:
			run = false
		default:
			return mem, fmt.Errorf("VM: invalid opcode [ip=%d]", ip)
		}
	}

	return mem, nil
}
