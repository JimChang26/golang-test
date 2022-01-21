package main

import (
	"fmt"
)

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle(req *Request) {
	req.resultChan <- req.f(req.args)
}

var clientRequests = make(chan *Request)

func main() {

	go func() {
		for {
			handle(<-clientRequests)
		}
	}()

	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// Send request
	clientRequests <- request
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultChan)
}
