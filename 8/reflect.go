package main

import (
	"fmt"
	"reflect"
)

//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
//通过反射机制，我们可以在运行时确定传入变量的类型，并根据类型执行不同的操作。这在某些需要处理未知类型数据的场景中非常有用，比如 JSON 解析、通用函数库的实现等。
/*
1.运行时类型检查和处理
	允许程序在运行时检查和操作未知类型的变量
	可以处理不同类型的数据而无需为每种类型写专门的代码
	特别适用于需要处理多种数据类型的通用代码
2.实现通用性和灵活性
	能够编写更加通用的函数和库
	可以处理在编译时无法确定类型的情况
	对于需要处理动态类型的场景非常有用
3.实际应用场景
	JSON/XML 等数据格式的序列化和反序列化
	ORM（对象关系映射）框架
	配置文件解析
	插件系统开发
	单元测试框架
*/

func IsNil_IsVaid() {
	// *int类型空指针
	var a *int
	//报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	//返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
