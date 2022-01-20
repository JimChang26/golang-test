package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("safely do failed ->", err)
		}
	}()
	panic("goroutine is dead")
}
