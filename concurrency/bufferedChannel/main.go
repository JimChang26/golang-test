package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	go func() {
		//time.Sleep(2 * time.Second)
		fmt.Println("receive channel data ", <-c)
		fmt.Println("receive channel data ", <-c)
		// fmt.Println("receive channel data ", <-c)
	}()
	data := 666
	fmt.Println("send channel data -> ", data)
	c <- data
	c <- data + 1
	c <- data + 2
	c <- data + 3
	c <- data + 4
	fmt.Println("send finished")
	// time.Sleep(10 * time.Second)
	// fmt.Println("receive channel data ", <-c)
	// fmt.Println("receive channel data ", <-c)
	// fmt.Println("receive channel data ", <-c)
}
