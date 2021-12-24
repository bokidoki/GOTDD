package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	stop chan struct{} = make(chan struct{})
)

func main() {
	// fmt.Println("start main")
	// go dosth()
	// <-stop
	// fmt.Println("get sth back")
	// fmt.Println("main done")
	for i := 0; i < 5; i++ {
		defer fmt.Println(strconv.Itoa(i))
	}
}

func dosth() {
	fmt.Println("start do sth")
	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	fmt.Println("do sth done")
}
