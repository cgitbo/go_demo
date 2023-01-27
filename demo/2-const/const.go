package main

import "fmt"

// const 定义枚举
// iota 自动递增
// iota 公式默认会使用上一个
const (
	BEIJING  = iota // 0
	SHANGHAI        // 1
	SHENZHEN        // 2
)

const (
	a, b = iota, iota + 2     // iota = 0, a = 0,     b = 0 + 2
	c, d                      // iota = 1, c = 1,     d = 1 + 2
	e, f = iota * 2, iota * 3 // iota = 2, e = 2 * 2, f = 2 * 3
	g, h = iota + 2, iota + 3 // iota = 3, g = 3 + 2, h = 3 + 3
	i, j                      // iota = 4, i = 4 + 2, j = 4 + 3
)

func main() {
	fmt.Println(BEIJING, SHANGHAI, SHENZHEN)
}
