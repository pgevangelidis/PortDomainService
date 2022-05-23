package parser

import (
	"fmt"
	"runtime"
	"testing"
)

func TestStreamJson(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// add tests
	printMemStats(m)
}

func BenchmarkStreamJson(b *testing.B) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	for i := 0; i < b.N; i++ {
		// add benchmark tests for memory stats
	}
	printMemStats(m)
}

func toMB(b uint64) uint64 {
	return b / 1024 / 1024
}

func printMemStats(m runtime.MemStats) {
	fmt.Printf("Alloc = %v MiB\n", toMB(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", toMB(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", toMB(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
