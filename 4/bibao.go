package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func dispatchCoin() int {
	totolcoins := coins

	rule := map[rune]int{
		'e': 1, 'E': 1,
		'i': 2, 'I': 2,
		'o': 3, 'O': 3,
		'u': 4, 'U': 4,
	}

	for _, user := range users {
		usercoins := 0
		for _, userchar := range user {
			if value, exists := rule[userchar]; exists {
				usercoins += value
			}
		}

		fmt.Println(usercoins)
		distribrtion[user] = usercoins
		totolcoins -= usercoins
	}

	return totolcoins
}
