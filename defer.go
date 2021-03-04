package main

import "fmt"

func d1() {
	for i:=3; i>0; i-- {
		defer fmt.Print(i, " ")
	}
	// 1. Step: fmt.Print(3, " ")
	// 2. Step: fmt.Print(2, " ")
	// 3. Step: fmt.Print(1, " ")
	// For LIFO, output: 3 2 1
}

func d2() {
	for i:= 3; i >0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	// defer for bittikten sonra yani i 0 olduktan sonra
	// çalıştığı için 0 0 0 yazıyor. d3deki gibi
	// parametre almamız gerek.
	fmt.Println()
}

func d3() {
	for i:= 3; i >0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
}
