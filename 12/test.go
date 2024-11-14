package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	count = 24
	wg    sync.WaitGroup
)

func mak(ch chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		num := rand.Int63()
		ch <- int(num)
	}
	close(ch)
}

func work(jobChan chan int, resultChan chan int) {
	defer wg.Done()

	for num := range jobChan {
		sum := 0
		n := num
		fmt.Println(n)
		for n > 0 {
			sum += int(n % 10)
			n /= 10
		}
		resultChan <- sum
	}
}

func prit(resultChan chan int) {
	for i := 0; i < count; i++ {
		result := <-resultChan
		fmt.Printf("第%d个数的各位数字之和为：%d\n", i+1, result)
	}
	close(resultChan)
}

func freeadd() {

	jobChan := make(chan int, 24)
	resultChan := make(chan int, 24)
	go mak(jobChan)
	wg.Add(1)
	for i := 0; i < count; i++ {
		go work(jobChan, resultChan)
		wg.Add(1)
	}
	prit(resultChan)
	wg.Wait()

}
