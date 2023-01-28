package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Printf("%v\n", u)
}

func reflectNum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func main() {
	num := 1.23456

	reflectNum(num)

	u := User{
		Id:   1,
		Name: "John",
		Age:  20,
	}

	u.Call()

	DoFileAndMethod(u)
}

func DoFileAndMethod(input interface{}) {
	// 获取type
	inputType := reflect.TypeOf(input)
	// 获取value
	inputValue := reflect.ValueOf(input)

	fmt.Println(inputType.Name(), inputValue)

	// 1.获取interface的reflect.Type 通过type得到NumField 进行遍历
	// 2.得到每个field 数据类型
	// 3.通过field有个interface()方法得到对应的value

	fmt.Println("================================================================")
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field, value)
		fmt.Printf("%s: %v => %v\n", field.Name, field.Type, value)
	}

	fmt.Println("================================")

	// 通过type 获取方法 调用
	fmt.Println(inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
