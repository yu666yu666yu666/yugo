package main

import "fmt"

func main() {
	mapone := make(map[string]int, 7)
	mapone["kdf"] = 33
	mapone["dgk"] = 95
	fmt.Println(mapone)
	fmt.Println(mapone["kdf"])

}
