package main

import "fmt"

func main() {
	// 声明slice1 是个切片，并且初始化，默认值 1，2，3 长度len为3
	// slice1 := []int{1, 2, 3}

	// 声明slice1 没分配空间
	var slice1 []int
	// slice1 = make([]int, 3) // 分配3个空间，初始化值为0

	// 常用方式
	// slice1 := make([]int, 3)

	// 判断slice 是否为0
	if slice1 == nil {
		println("slice1 is nil")
	}

	fmt.Printf("%d, v=%v\n", len(slice1), slice1)

	// 数组截取 用的是地址引用，如果要用原值，需要用 copy 函数
	s := []int{1, 2, 3}
	fmt.Println("s:", s)

	s1 := s[0:2]
	fmt.Println("s1:", s1)

	s[0] = 999

	fmt.Println("s:", s)
	fmt.Println("s1:", s1)

	// 深拷贝后 原数组改变，新数组不再改变
	s2 := make([]int, 3)
	copy(s2, s)

	s[0] = 666

	fmt.Println("s,s1,s2:", s, s1, s2)
}
