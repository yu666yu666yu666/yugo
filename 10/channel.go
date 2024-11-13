package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main1() {
	ch := make(chan int)     //无缓冲
	ch2 := make(chan int, 3) //有缓冲
	go recv(ch)              // 创建一个 goroutine 从通道接收值
	ch <- 10
	ch2 <- 10
	fmt.Println("发送成功")
	//value, ok := <-ch	多返值模式可以判断是否取完

}

// 通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

//**注意：**目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过len(ch)操作来判断通道是否被关闭。
//在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。
