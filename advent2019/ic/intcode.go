package ic

import (
	"fmt"
	"io"
)

// Options is used to keep future changes from break the VM.
type Options struct {
	Console io.ReadWriter
}

func DefaultOptions() *Options {
	return &Options{
		Console: Console(),
	}
}

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
	opOutput: 1,
	opJT:     2,
	opJF:     2,
	opLT:     2,
	opEQ:     2,
	opHalt:   0,
}

func fetchArgs(mem []int, ip, params, count int) []int {
	if count == 0 {
		return nil
	}

	args := make([]int, 0, count)
	for i := 1; i <= count; i++ {
		arg := mem[ip+i]
		if params%10 == 0 {
			arg = mem[arg]
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
func Run(prog []int, opts *Options, mods ...mod) ([]int, error) {
	if opts == nil {
		opts = DefaultOptions()
	}

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
		args := fetchArgs(mem, ip, params, argCount[op%100])
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
			var arg int
			fmt.Fprintf(opts.Console, "? ")
			fmt.Fscan(opts.Console, &arg)

			dest := mem[ip+1]
			mem[dest] = arg
			ip += 2
		case opOutput:
			fmt.Fprintln(opts.Console, args[0])
			ip += 2
		case opJT:
			if args[0] != 0 {
				ip = args[1]
			} else {
				ip += 3
			}
		case opJF:
			if args[0] == 0 {
				ip = args[1]
			} else {
				ip += 3
			}
		case opLT:
			dest := mem[ip+3]
			if args[0] < args[1] {
				mem[dest] = 1
			} else {
				mem[dest] = 0
			}
			ip += 4
		case opEQ:
			dest := mem[ip+3]
			if args[0] == args[1] {
				mem[dest] = 1
			} else {
				mem[dest] = 0
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

func Dump(mem []int, w io.Writer) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "\t%05d", i)
	}
	fmt.Fprintln(w)
	i := 0
	fmt.Fprintf(w, "%05d", i)
	for i < len(mem) {
		fmt.Fprintf(w, "\t%05d", mem[i])
		i++
		if i%10 == 0 {
			fmt.Fprintf(w, "\n%05d", i)
		}
	}

	if len(mem)%10 != 0 {
		fmt.Fprintln(w)
	}
}
