package main

import (
	"fmt"
)

var qp []int

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
}
