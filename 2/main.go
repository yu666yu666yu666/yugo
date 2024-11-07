package main

import (
	"fmt"
)

var qp = make([]string, 5, 10)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d    ", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println()

	array1 := []int{1, 3, 5, 7, 8}

	array_special_sum(array1, 8)

	a := make([]int, 3, 5)

	fmt.Println(a)

	b := []int{1, 3, 4, 5, 6}
	b = append(b, 3, 4)
	c := b[1:3]
	fmt.Println(b)
	fmt.Println(c)
	d := c[3:4]
	fmt.Println(d)

	//删除下标为2的元素
	b = append(b[:2], b[3:]...)
	fmt.Println(b)

	for i := 0; i < 10; i++ {
		qp = append(qp, fmt.Sprintf("%v", i))
		if qp[i] == "" {
			fmt.Println(1)
		}
	}
	fmt.Println(qp)
}
