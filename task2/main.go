package main

import (
	"fmt"
	"sync"
)

func main() {

	target := 6
	fmt.Printf("原数据：%d\n", target)
	addTen(&target)
	fmt.Printf("输出指针加10后的数字：%d\n", target)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("原数据：%d\n", arr)
	mutiply2(&arr)
	fmt.Printf("输出指针乘2后的数字：%d\n", arr)

	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber1(&wg)
	go printNumber2(&wg)
	wg.Wait()

}

// 指针加10
func addTen(num *int) {
	*num += 10
}

// 将切片中的每个元素乘以2
func mutiply2(arr *[]int) {

	for i, v := range *arr {
		(*arr)[i] = v * 2
	}

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
