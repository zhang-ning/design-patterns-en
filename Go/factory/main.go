package main

import (
	"fmt"
	"time"
)

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)

	// var quit = make(chan bool)

	// Serve(clientRequests, quit)
	// Serve(clientRequests, quit)

	go SendReq([]int{1, 2, 3})
	go SendReq([]int{4, 5, 6})
	go SendReq([]int{7, 8, 9})

	// }
}

func SendReq(args []int) <-chan *Request {
	var clientRequests = make(chan *Request)
	go func() {
		request := &Request{args, sum, make(chan int)}
		// Send request
		clientRequests <- request

		// Wait for response.
		fmt.Printf("answer: %d\n", <-request.resultChan)
	}()

	return clientRequests
}

var MaxOutstanding int = 3

// var sem = make(chan int, MaxOutstanding)

// func process() {
// 	time.Sleep(time.Second * 1)
// }

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	time.Sleep(time.Second * 1)
	return
}

func handle(queue chan *Request) {
	// time.Sleep(time.Second * 1)
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // Wait to be told to exit.

}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
