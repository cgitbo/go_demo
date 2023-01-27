package main

import "fmt"

func main() {
	r1, r2 := foo("foo", "bar")
	fmt.Println(r1, r2)
	var r3 string
	println(r3)
}

// 返回值可以写名称 也可以不写
func foo(a string, b string) (string, int) {
	return "1", 2
}

func foo1(a string, b string) (r1 string, r2 int) {
	println(r1, r2)
	return "1", 2
}
