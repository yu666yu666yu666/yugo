package main

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
