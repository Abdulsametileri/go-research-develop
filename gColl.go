package main

import (
	"fmt"
	"runtime"
)

func printStats(mem *runtime.MemStats) {
	runtime.ReadMemStats(mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("--------")
}

func main() {
	var mem runtime.MemStats
	printStats(&mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(&mem)
	// GODEBUG=gctrace=1 go run gColl.go - for a detailed output
	// OUTPUT: .... ... ... 47->47->0 MB
	// The first number is heap size GC tam çalışacakken
	// The second number is heap size GC çalışmayı bitirince
	// The third number is Live heap
}
