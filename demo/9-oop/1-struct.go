package main

import "fmt"

// 声明数据类型 myint 是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	// 传递的是一个副本 这里改变 并不会改变原结构体
	book.title = "Book"
}

func changeBook2(book *Book) {
	// 指针传递
	book.title = "Book 2"
}

func main() {

	var book1 Book

	book1.title = "Golang"
	book1.author = "zhang san"

	fmt.Println(book1)
	// 不改变原结构体
	changeBook(book1)

	fmt.Println(book1)

	// 指针传递，可以改变原结构体
	changeBook2(&book1)
	fmt.Println(book1)
}
