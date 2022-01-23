package main

import "fmt"

func main() {
	var a interface{} = 10.00
	fmt.Println(a.(float32))
}
