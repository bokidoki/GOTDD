package main

import (
	"flag"
	"fmt"
	"os"
)

// "fmt"
// "os"

var a string

func main() {
	// fileinfo, err := os.Stat(filepath.Join("variables.go"))
	// fmt.Printf("fileinfo is %#v, err is %s", fileinfo, err)

	// args := os.Args
	// fmt.Printf("args are %v", args)

	// args := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("从下标2开始截取args得到的数组为：%v", args[2:])

	dlgo := flag.Bool("dlgo", false, "Download Go and build with it")

	flag.CommandLine.Parse(os.Args[:2])

	fmt.Printf("dlgo is %v", *dlgo)

}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

func n() { print(a) }

func m() {
	a = "O"
	print(a)
}
