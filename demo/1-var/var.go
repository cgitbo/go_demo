package main

// var 变量
// 方法一 指定变量类型，声明后若不赋值，使用默认值0。
var a int
var a1 string = "11"

// 方法二 根据值自行判定变量类型。
var b = 1

func main() {
	// 方法三 只能用于函数内使用 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	c := "1"

	var _, xx, yy = 0, "0", true
	println(a, b, c, xx, yy)
}
