package main

import (
	"fmt"
	"sync"
	"time"
)

func routine() {
	for {
		select {
		case <-pause:
			fmt.Println("pause")
			// select {
			// case <-play:
			// 	fmt.Println("play")
			// case <-quit:
			// 	fmt.Println("quit inside")
			// 	wg.Done()
			// 	return

			// }
		case <-pause2:
			fmt.Println("pause2")
		case <-quit:
			fmt.Println("quit outside")
			wg.Done()
			return
		default:
			work()
		}
	}
}

func main() {
	wg.Add(1)
	go routine()

	time.Sleep(1 * time.Second)
	pause <- struct{}{}
	fmt.Println("send pause channel")
	pause2 <- struct{}{}
	fmt.Println("send pause2 channel")
	time.Sleep(1 * time.Second)
	//play <- struct{}{}

	time.Sleep(1 * time.Second)
	pause <- struct{}{}

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	close(quit)

	wg.Wait()
	fmt.Println("done")
}

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Print(i, " ")
}

var play = make(chan struct{})
var pause = make(chan struct{})
var pause2 = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 0
