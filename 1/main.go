package main

import (
	"fmt"
	"strings"
	"unicode"
)

func foo() (int, string) {
	return 10, "yu"
}

var (
	a    string
	b    int
	c    bool
	cc   = true
	d    float32
	dd   = 6.6
	age  = 18
	name = "yu"
)

func main() {
	fmt.Println("Hello World!")
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	s1 := `
	sdf
	sdf
	`
	fmt.Println(s1)
	fmt.Println(len(name))
	too := "a,d,g"
	boo := strings.Split(too, ",")
	fmt.Println(boo)

	s := "Hello, World"
	index := strings.Index(s, "World")
	lastIndex := strings.LastIndex(s, "o")
	fmt.Println(index)     // 输出：7
	fmt.Println(lastIndex) // 输出：8

	fmt.Printf("%T %T %T %T\n", cc, dd, age, name)

	//byte
	s2 := "白萝卜"
	runes2 := []rune(s2)
	runes2[0] = '红'
	fmt.Println(string(runes2))

	bb := "hello沙河小王子"
	fmt.Println(bb)
	bbb := len(bb)
	fmt.Println(bbb)
	count := 0
	for _, rrr := range bb {
		if unicode.Is(unicode.Han, rrr) {
			count++
		}
	}
	fmt.Println(count)

	ooo := []int{1, 1, 546, 45, 45}
	fmt.Println(ooo)
	pp := 0
	for _, o := range ooo {
		pp ^= o
	}
	fmt.Println(pp)

	switch1()
}

const (
	pi = 3.1415
	e  = 1.2
)
