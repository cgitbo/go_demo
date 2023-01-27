package main

import (
	_ "fmt" // 导入 不使用，但是会执行 init
	f "fmt" // 导入 并起别名
	// "lib1"
	// "lib2"
)

func main() {
	f.Println('1')
	// lib1.Lib1Test()
	// lib2.Lib1Test()
}
