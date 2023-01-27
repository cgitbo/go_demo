package main

import "fmt"

func main() {
	// 动态数组 长度不固定
	myArr := []int{1, 2, 3}

	println("===")
	for _, v := range myArr {
		println(v)
	}

	printArray(myArr)

	println("===")
	for _, v := range myArr {
		println(v)
	}

}

func printArray(myArr []int) {
	fmt.Printf("%T\n", myArr)
	// 引用值传递
	for _, v := range myArr {
		println(v)
	}
	myArr[0] = 666

}
