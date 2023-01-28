package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex" doc:"性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfoStr := t.Field(i).Tag.Get("info")
		fmt.Println(tagInfoStr)

		tagDocStr := t.Field(i).Tag.Get("doc")
		fmt.Println(tagDocStr)
	}
}

func main() {
	var re resume

	// 需要传引用值
	findTag(&re)
}
