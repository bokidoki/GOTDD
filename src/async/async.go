package main

import (
	"fmt"
	"time"
)

var (
	stop chan struct{} = make(chan struct{})
)

func main() {
	broad := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			broad <- i
		}
	}()

	go func() {
		for {
			recv := <-broad
			time.Sleep(time.Second * 3)
			fmt.Println(recv)
		}
	}()

	time.Sleep(time.Minute)
}

func dosth() {
	fmt.Println("start do sth")
	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	fmt.Println("do sth done")
}
