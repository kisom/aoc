// Package inst contains tracing and monitoring code for instrumenting
// code.
package inst

import (
	"log"
	"runtime"
	"time"
)

// Run executes function f, printing total runtime and allocation
// sizes at the end.
func Run(label string, f func()) {
	var start, end int64
	var mem runtime.MemStats

	start = time.Now().UnixNano()
	f()
	end = time.Now().UnixNano()

	runtime.ReadMemStats(&mem)
	runtime := (end - start) / 1000
	tAlloc := mem.TotalAlloc / 1024
	alloc := mem.HeapAlloc / 1024
	log.Printf("%s: complete in %dus, total allocated %d kB, current heap allocation %d kB",
		label, runtime, tAlloc, alloc)
}
