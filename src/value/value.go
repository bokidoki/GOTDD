package value

import "fmt"

// Aboutable ...
type Aboutable interface {
	About() string
	NotAbout() string
}

// Book ...
type Book struct {
	name string
}

// About ...
func (book *Book) About() string {
	return "Book:" + book.name
}

func (book *Book) NotAbout() string {
	return "nothing to say"
}

type Filter interface {
	About() string
	Process([]int) []int
}

type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "删除重复的数字"
}

func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("保留%v的倍数", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

func filteAndPrint(fltr Filter, unfiltered []int) []int {
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}
