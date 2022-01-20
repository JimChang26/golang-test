package main

import (
	"fmt"
	"sync"
)

func main() {

	var x interface{} = 2

	fmt.Print(x)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go safelyDo(wg)
	wg.Wait()
}

func safelyDo(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("safely do failed ->", err)
		}
		wg.Done()
	}()
	do()
}

func do() {
	for i := 0; i < 100; i++ {
		fmt.Println("do!")
		if i > 10 {
			panic("goroutine is dead")
		}
	}
}
