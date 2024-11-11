package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func mapsort() {

	rand.Seed(time.Now().UnixNano())
	var map1 = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		map1[key] = value
	}

	keys := make([]string, 0, 200)
	for key := range map1 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, map1[key])
	}
}

func numword() {
	input := "how do you do"
	words := strings.Fields(input)
	wordcount := make(map[string]int)
	for _, word := range words {
		wordcount[word]++
	}
	for word, count := range wordcount {
		fmt.Println(word, count)
	}

}

// 在处理切片时要特别注意，如果不想互相影响，应该使用copy函数创建新的切片。
func examp() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
