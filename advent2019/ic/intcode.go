package ic

import (
	"fmt"
	"io"
)

const (
	opAdd    = 1
	opMul    = 2
	opInput  = 3
	opOutput = 4
	opJT     = 5
	opJF     = 6
	opLT     = 7
	opEQ     = 8
	opHalt   = 99
)

var argCount = map[int]int{
	opAdd:    2,
	opMul:    2,
	opInput:  0,
	opOutput: 0,
	opJT:     2,
	opJF:     2,
	opLT:     2,
	opEQ:     2,
	opHalt:   0,
}

type mod struct {
	Pos int
	Val int
}

// Mods encode changes to the program to be made at runtime.
func Mod(pos, val int) mod {
	return mod{
		Pos: pos,
		Val: val,
	}
}

func fetchArgs(mem *[]int, ip, params, count int) []int {
	args := make([]int, 0, count)

	for i := 1; i <= count; i++ {
		arg := (*mem)[ip+i]
		if params%10 == 0 {
			arg = (*mem)[arg]
		}
		params /= 10

		args = append(args, arg)
	}

	return args
}

// Run copies the program to a new memory space, applies any mods, and
// the runs it, returning the program. For Day 2, all of the results
// are in memory position 0, but I don't want to make that assumption
// for the future.
func Run(prog []int, r io.Reader, mods ...mod) ([]int, error) {
	mem := make([]int, len(prog))
	copy(mem, prog)

	for _, mod := range mods {
		mem[mod.Pos] = mod.Val
	}

	ip := 0
	run := true
	for run {
		op := mem[ip] % 100
		params := mem[ip] / 100
		args := fetchArgs(&mem, ip, params, argCount[op%100])
		switch op % 100 {
		case opAdd:
			dest := mem[ip+3]
			mem[dest] = args[0] + args[1]
			ip += 4
		case opMul:
			dest := mem[ip+3]
			mem[dest] = args[0] * args[1]
			ip += 4
		case opInput:
			fmt.Printf("? ")
			var arg int
			fmt.Fscan(r, &arg)
			dest := mem[ip+1]
			mem[dest] = arg
			ip += 2
		case opOutput:
			dest := mem[ip+1]
			fmt.Printf("%08d: %d\n", dest, mem[dest])
			ip += 2
		case opJT:
			if args[0] != 0 {
				ip = args[1]
			} else {
				ip += 2
			}
		case opJF:
			if args[0] == 0 {
				ip = args[1]
			} else {
				ip += 2
			}
		case opLT:
			if args[0] < args[1] {
				mem[ip+3] = 1
			} else {
				mem[ip+3] = 0
			}
			ip += 4
		case opEQ:
			if args[0] == args[1] {
				mem[ip+3] = 1
			} else {
				mem[ip+3] = 0
			}
			ip += 4
		case opHalt:
			run = false
		default:
			return mem, fmt.Errorf("VM: invalid opcode %d [ip=%d]", op, ip)
		}
	}

	return mem, nil
}
