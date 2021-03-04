package main

import (
	"fmt"
	"log"
	"os"
)

func one(aLog *log.Logger) {
	aLog.Println("-- FUNCTION ONE ------")
	defer aLog.Println("-- FUNCTION ONE ------")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile("/tmp/mGo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
	one(iLog)
}
