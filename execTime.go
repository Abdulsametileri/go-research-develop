package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	duration := time.Since(start)
	fmt.Println("It took time.Sleep(2)", duration, "to finish")
}
