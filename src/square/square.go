package square

import (
	"math"
)

// Shape ...
// go 中的接口定义和其它编程语言有所区别
// go 中的实现是隐式的。两个类型之间的实现关系不需要在代码中显式地表现出来
// go 中没有类似于implements的关键字。go 编译器将自动在需要的时候检查两个类型之间的实现关系。
// 任何方法集都是空方法集的超集，所以任何类型都实现了任何空接口类型。
type Shape interface {
	Area() float64
}

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Area ...
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
// 声明方法 func(receiverName ReceiverType) MethodName(args)
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Triangle ...
type Triangle struct {
	Base   float64
	Height float64
}

// Area ...
func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height / 2
}

// Perimeter ...
// 声明函数
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}
