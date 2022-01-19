package main

import (
	"fmt"
	"time"
)

var (
	stop chan struct{} = make(chan struct{})
)

func main() {
	// doselect(make(chan int), make(chan int))
	testDefer()
}

func dosth() {
	fmt.Println("start do sth")
	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	fmt.Println("do sth done")
}

func doselect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x)
			fmt.Println(y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
}
