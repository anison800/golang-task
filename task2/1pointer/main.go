package main

import (
	"fmt"
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
