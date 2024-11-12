package main

import "fmt"

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams1(dreams []string) { //错误
	p.dreams = dreams
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams1(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?
}

func (p *Person) SetDreams2(dreams []string) { //正确
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

/*
切片共享底层数组的特性：
切片是对底层数组的引用
直接赋值切片会导致多个切片指向同一个底层数组
深拷贝的必要性：
在需要独立数据的场景中，应该使用深拷贝
make + copy 的组合可以实现切片的深拷贝
安全性考虑：
在设计方法时要考虑数据的独立性
防止外部修改影响内部数据
其他可能的解决方案：

// 方案1：使用append
func (p *Person) SetDreams(dreams []string) {
    p.dreams = append([]string{}, dreams...)
}

// 方案2：手动遍历复制
func (p *Person) SetDreams(dreams []string) {
    p.dreams = make([]string, len(dreams))
    for i, v := range dreams {
        p.dreams[i] = v
    }
}
*/
