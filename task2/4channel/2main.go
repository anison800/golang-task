package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(2)
	go sendMsg2(ch, &wg)
	go recieveMsg2(ch, &wg)
	wg.Wait()

	var input string
	fmt.Scanln(&input)
	fmt.Println("程序结束")
}

func sendMsg2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func recieveMsg2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("输出：", val)
	}
}
