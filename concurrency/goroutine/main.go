package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("child goroutine")
	}()
	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
