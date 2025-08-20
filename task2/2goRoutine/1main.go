package main

import (
	"fmt"
	"sync"
)

func main() {

	//1、打印奇数偶数
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber1(&wg)
	go printNumber2(&wg)
	wg.Wait()

}

// 打印奇数
func printNumber1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("奇数：%d ", i)
		}
	}
}

// 打印偶数
func printNumber2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数：%d ", i)
		}
	}
}
