package main

import "fmt"

func main() {
	// 初始化map
	var myMap1 map[string]string
	if myMap1 == nil {

		println("mymap is nil")
	}

	//在使用之前，先分配空间
	myMap1 = make(map[string]string, 3)

	myMap1["one"] = "one"
	myMap1["two"] = "two"

	fmt.Println(myMap1)

	// 第二种方式
	myMap2 := make(map[int]string)
	myMap2[1] = "one"
	myMap2[2] = "two"
	fmt.Println(myMap2)

	// 第三种 声明并赋值
	myMap3 := map[string]string{
		"one":   "PHP",
		"two":   "Python",
		"three": "GO",
	}
	fmt.Println(myMap3)

	// 遍历
	printMap(myMap3)

	// 删除
	delete(myMap3, "one")

	fmt.Println("-----")

	// 修改
	myMap3["ABC"] = "ABC"
	printMap(myMap3)
}

// 引用值
// 如果需要深拷贝，需要定义个新map 遍历赋值
func printMap(myMap map[string]string) {
	// 遍历
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
