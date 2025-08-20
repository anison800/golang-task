package main

import "fmt"

func main() {

	ch := make(chan int)

	go sendMsg(ch)
	go reciveMsg(ch)

	var input string
	fmt.Scanln(&input) // 等待用户输入，防止主程序提前退出
	fmt.Println("程序结束")
}

func sendMsg(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func reciveMsg(ch chan int) {
	for val := range ch {
		fmt.Println("输出接受数据:", val)
	}
}
