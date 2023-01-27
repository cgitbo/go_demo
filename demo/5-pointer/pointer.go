package main

func main() {
	a := 10
	b := 20
	swap(&a, &b)

	println(a, b)

	println(&a, &b)
}

// 地址引用赋值 行参用 * 实参用 &
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	println(a, b) // 这里使用的是 地址
}
