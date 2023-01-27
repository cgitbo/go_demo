package main

import "fmt"

// 万能数据类型 interface{}
func myFunc(arg interface{}) {
	fmt.Println(arg)

	// 类型断言 判断是否字符串
	v, ok := arg.(string)
	fmt.Println(v, ok)
}

type Book struct {
	author string
}

func main() {
	book := Book{"Go"}

	myFunc(book)
	myFunc(123)
	myFunc("666")
	myFunc(1.22)
	myFunc(false)
}
