package main

/*
#include <stdio.h>
void callC() {
    printf("Calling C code!\n");
}
*/
import "C"
// there should be no newline between the /* ... */ cgo block and import "C"
import "fmt"

func main() {
	fmt.Println("A Go Statement")
	C.callC()
	fmt.Println("Another Go Statement")
}
