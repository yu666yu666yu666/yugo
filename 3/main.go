package main

import "fmt"

func main() {
	mapone := make(map[string]int, 7)
	mapone["kdf"] = 33
	mapone["dgk"] = 95
	fmt.Println(mapone)
	fmt.Println(mapone["kdf"])
	iff, ok := mapone["kdf"]
	if ok {
		fmt.Println(iff)
	} else {
		fmt.Println("wu")
	}

	mapsort()
	examp()
	numword()
}
