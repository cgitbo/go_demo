package main

func main() {
	fun()
}

func fun() {
	// defer 相当于 final
	// 当前函数生命周期结束后执行
	// 在return后执行
	// 先进后出 FILO
	defer defer1()
	defer defer2()
	defer defer3()
	returnFun()
}

func defer1() {
	println("defer1")
}

func defer2() {
	println("defer2")
}

func defer3() {
	println("defer3")
}

func returnFun() {
	println("returnFun")
}
