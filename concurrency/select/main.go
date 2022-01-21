package main

import (
	"fmt"
	"time"
)

var channelA = make(chan int)
var channelB = make(chan int)

func routine() {
	for {
		select {
		case a := <-channelA:
			fmt.Println("channelA get", a)

		case b := <-channelB:
			fmt.Println("channelB get", b)
			// default:
			// 	fmt.Println("default")
		}

		fmt.Println("for loop")
	}
}

func main() {
	go routine()

	go func() {
		channelA <- 1
	}()
	go func() {
		channelB <- 1
	}()
	time.Sleep(2 * time.Second)
}
