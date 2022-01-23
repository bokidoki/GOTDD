package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("do wait for work")
			time.Sleep(time.Second)
			defer wg.Done()
			fmt.Println("work done")
		}()
	}

	wg.Wait()
	fmt.Println("all works done")
}
