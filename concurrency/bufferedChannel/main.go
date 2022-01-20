package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("receive channel data")
		<-c
	}()
	data := 666
	fmt.Println("send channel data -> ", data)
	c <- data
	c <- data
	fmt.Println("send finished")
}
