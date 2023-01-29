package main

import (
	"fmt"
	"time"
)

func main() {
	// 管道通信
	c := make(chan int)

	go func() {
		// 写入文件
		c <- 666
		fmt.Println("running")
		defer fmt.Println("end")
	}()

	go func() {
		fmt.Println("running2")
		c <- 1
		defer fmt.Println("end2")
	}()

	// 读取
	for {
		time.Sleep(time.Second)
		fmt.Println(<-c)
	}

}
