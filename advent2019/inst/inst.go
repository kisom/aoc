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
func Run(f func()) {
	var start, end time.Time
	var mem runtime.MemStats

	start = time.Now()
	f()
	end = time.Now()

	runtime.ReadMemStats(&mem)
	runtime := end.Sub(start).Milliseconds()
	tAlloc := mem.TotalAlloc / 1024
	alloc := mem.Alloc / 1024
	log.Printf("complete in %dms, total allocated %d kB, current heap allocation %d kB",
		runtime, tAlloc, alloc)
}
