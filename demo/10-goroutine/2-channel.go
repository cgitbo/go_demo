package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("end")
		fmt.Println("running")
	}()

	fmt.Println(c)
}
