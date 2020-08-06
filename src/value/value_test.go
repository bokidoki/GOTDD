package value

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a.NotAbout())

	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i)

	i = a
	fmt.Println(i)
}

func TestPolymorphism(t *testing.T) {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("过滤之前：\n\t", numbers)
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}
	for _, fltr := range filters {
		numbers = filteAndPrint(fltr, numbers)
	}
}
