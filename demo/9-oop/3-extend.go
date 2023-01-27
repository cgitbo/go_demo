package main

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	println("h.eat...")
}

func (h *Human) Walk() {
	println("h.walk...")
}

// 子类继承
type SuperMan struct {
	Human // SuperMan 继承了 Human类 只需要把父类的类名写进来就可以
	level int
}

// 子类复写父类方法
func (h *SuperMan) Eat() {
	println("superman.eat..")
}

// 子类新方法
func (h *SuperMan) Fly() {
	println("superman.fly...")
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	println("====")
	// 第一种使用方式 初始化时候直接初始化值
	// superMan := SuperMan{Human{"spider", "male"}, 99}

	// 第二种方式，初始化后，再赋值 可读性更高
	var superMan SuperMan
	superMan.name = "spider"
	superMan.sex = "male"
	superMan.level = 100

	superMan.Walk() // 父类方法
	superMan.Fly()  // 子类方法
	superMan.Eat()  // 子类方法
}
