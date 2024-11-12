package main

import "fmt"

// 使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func a() {
	// 使用空接口实现可以保存任意值的字典。
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

//可以一个类型多接口，可以多个类型一接口，接口可以嵌套接口，接口可以做为结构体一个字段

/*
func b() {

	var n Mover = &Dog{Name: "旺财"}
	//第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}

}
*/

// justifyType 对传入的空接口类型变量x进行类型断言...如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
