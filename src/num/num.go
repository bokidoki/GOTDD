package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mode uint32

const (
	Mode1 Mode = iota
	Mode2
	Mode3
)

func main() {
	// fmt.Printf("mode1 is %v, mode2 is %v, mode3 is %v", Mode1, Mode2, Mode3)
	s := strconv.Itoa(2098)
	type_s := reflect.TypeOf(s)
	fmt.Printf("s is %s", type_s)
}
