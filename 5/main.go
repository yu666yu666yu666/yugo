package main

import "fmt"

//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

type Animal struct {
	name string
}

type Dog struct {
	feet int
	*Animal
}

func (a *Animal) move() {
	fmt.Println("会动")
}

func (d *Dog) wang() {
	fmt.Println("会汪")
}

func main() {
	d1 := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	d1.move()
	d1.wang()
}
