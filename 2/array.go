package main

import "fmt"

func array_special_sum(arr []int, target int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Printf("%d %d	", i, j)
			}
		}
	}
}
