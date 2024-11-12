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

	rules := map[rune]int{
		'e': 1, 'E': 1,
		'i': 2, 'I': 2,
		'o': 3, 'O': 3,
		'u': 4, 'U': 4,
	}

	//rules[letter]：尝试从 rules 映射中获取 letter 对应的值。
	//value：如果 letter 存在于 rules 中，则 value 将赋值为对应的金币数。
	//exists：这是一个布尔值，表示 letter 是否在 rules 中存在：
	//	如果存在，exists 为 true，并且 value 将是对应的金币数。
	//	如果不存在，exists 为 false，value 将为该类型的零值（在此情况下为 0）。
	for _, user := range users {
		usercoins := 0
		for _, letter := range user {
			if value, exists := rules[letter]; exists {
				usercoins += value
			}
		}

		fmt.Println(usercoins)
		distribrtion[user] = usercoins
		totolcoins -= usercoins
	}

	return totolcoins
}
