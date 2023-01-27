package main

import (
	"fmt"
	"math"
)

// 接口，多个类使用同一接口，必须实现这些方法
// 基本要素 1.父类（接口）2.子类（实现父类的全部接口方法）3.调用父类方法实际上实现的是自类的方法
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	fmt.Println("r.area")
	return r.width * r.height
}
func (r rect) perim() float64 {
	fmt.Println("r.perim")
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	fmt.Println("c.area")
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	fmt.Println("c.perim")
	return 2 * math.Pi * c.radius
}

// 多态 不同的类有同样的方法，一起调用
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
