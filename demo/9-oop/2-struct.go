package main

import "fmt"

// 如果类名首字母大写 表示其他包也可以访问 否则只能本包访问
type Hero struct {
	Name  string
	Ad    int
	Level int
}

// 获取名称
// 类方法首字母大写 代表public，如果小写 代表 private 只能本包使用
func (this *Hero) GetName() string {
	return this.Name
}

// 设置名称 要改对象内容 需要引用 加 *
func (this *Hero) SetName(name string) {
	this.Name = name
}

// 显示
func (this *Hero) Show() {
	fmt.Println(this)
}

func main() {

	hero := Hero{
		Name:  "hero",
		Ad:    100,
		Level: 99,
	}
	hero.Show()

	hero.SetName("new Name")
	hero.Show()
}
