package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 1
	m["bar"] = 2
	fmt.Printf("m is %v", m)

	var n map[string]int = map[string]int{
		"bar": 1,
		"foo": 2,
	}
	n["nicer"] = 3
	fmt.Printf("n is %v", n)

}
