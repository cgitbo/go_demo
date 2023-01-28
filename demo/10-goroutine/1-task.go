package main

import (
	"fmt"
	"runtime"
	"time"
)

func toTasks() {
	i := 0
	for {
		fmt.Printf("Task %d\n", i)
		time.Sleep(time.Second)
		i++

		if i == 10 {
			fmt.Printf("Task out %d\n", i)
			runtime.Goexit()
		}
	}
}
func toTasks1() {
	i := 0
	for {
		fmt.Printf("Task1 %d\n", i)
		time.Sleep(time.Second)
		i++
		if i == 11 {
			fmt.Printf("Task1 out %d\n", i)
			runtime.Goexit()
		}
	}
}
func toTasks2() {
	i := 0
	for {
		fmt.Printf("Task2 %d\n", i)
		time.Sleep(time.Second)
		i++
		if i == 12 {
			runtime.Goexit()
			fmt.Printf("Task2 out %d\n", i)
		}
	}
}
func toTasks3() {
	i := 0
	for {
		fmt.Printf("Task3 %d\n", i)
		time.Sleep(time.Second)
		i++
		if i == 13 {
			fmt.Printf("Task3 out %d\n", i)
			runtime.Goexit()
		}
	}
}
func toTasks4() {
	i := 0
	for {
		fmt.Printf("Task4 %d\n", i)
		time.Sleep(time.Second)
		i++
		if i == 14 {
			fmt.Printf("Task4 out %d\n", i)
			runtime.Goexit()
		}
	}
}

func main() {
	fmt.Println("--")

	// 创建go程 执行task方法
	go toTasks()
	go toTasks1()
	go toTasks2()
	go toTasks3()
	go toTasks4()

	// 匿名方法
	go func() {
		defer fmt.Println("a.defer")

		func() {
			defer fmt.Println("b.defer")
			runtime.Goexit()
		}()

		fmt.Println("A")
	}()

	// 主goroutine退出后，其它的工作goroutine也会自动退出
	for i := 0; i < 2; i++ {
		fmt.Printf("main %d\n", i)
		time.Sleep(time.Second)
	}
}
