package main

import "fmt"

func main() {
	// 固定长度数组 传递使用值拷贝 不改变原数组
	var myArr1 [10]int
	myArr2 := [6]int{1, 2, 3, 4}

	for i := 0; i < len(myArr1); i++ {
		println(myArr1[i])
	}

	for _, v := range myArr2 {
		println(v)
	}
	fmt.Printf("%T", myArr1)
	fmt.Printf("%T", myArr2)
}
