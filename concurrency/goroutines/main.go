package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10000; i++ {
		go printTime(i)
	}
	time.Sleep(3 * time.Second)
}

func printTime(count int) {
	fmt.Println(count, " ", time.Now())
}
