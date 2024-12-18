package main

import (
	"fmt"
	"sync"
)

// sync.Mutex

var (
	x1 int64

	wg1 sync.WaitGroup // 等待组

	m sync.Mutex // 互斥锁
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x1 = x1 + 1
		m.Unlock() // 改完解锁
	}
	wg1.Done()
}

func main() {
	wg1.Add(2)

	go add()
	go add()

	wg1.Wait()
	fmt.Println(x1)
}
