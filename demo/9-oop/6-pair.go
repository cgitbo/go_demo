package main

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	println("read book")
}

func (b *Book) WriteBook() {
	println("write book")
}

func main() {
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}

// 断言有两步：得到动态类型 type，判断 type 是否实现了目标接口。
// 这里断言成功是因为 type 是 Book，而 Book 实现了 Writer 接口
