package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		data := 666
		fmt.Println("send channel data ", data)
		c <- data
		fmt.Println("send finished")
	}()
	fmt.Println("receive channel data")
	i := <-c
	fmt.Println("revice ", i)
}
