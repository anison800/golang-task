package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 只出现一次的数字
func singleNum(list []int) int {
	mp := make(map[int]int)

	for _, num := range list {
		if mp[num] == 0 {
			mp[num] = 1
		} else {
			mp[num]++
		}
	}

	var result int
	for key, val := range mp {
		if val == 1 {
			result = key
			break
		}
	}
	fmt.Println("Hello, World!", list, mp, result)

	return result
}

// 只出现一次的数字2
func singleNum2(list []int) int {
	result := 0
	for _, val := range list {
		result = result ^ val
	}
	fmt.Printf("输出%d", result)
	return result
}

// 判断回文数
func isPalindrome(x int) bool {
	// 1. 负数不是回文数
	// 2. 将其倒序数值计算出来，然后比较和原数值是否相等
	if x < 0 {
		return false
	}

	cur := 0
	num := x
	for num != 0 {
		cur = cur*10 + num%10
		num = num / 10
	}

	fmt.Printf("是否是回文数：%t\n", cur == x)

	return cur == x

}

// 判断回文数
func isPalindrome2(x int) bool {

	x_str := strconv.Itoa(x)

	var rever_str string
	for i := len(x_str) - 1; i >= 0; i-- {
		rever_str = rever_str + string(x_str[i])
	}

	rever_num, _ := strconv.Atoi(rever_str)

	fmt.Printf("%d是否是回文数：%t\n", x, x == rever_num)

	return x == rever_num

}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(str string) bool {
	stack := []rune{}

	mp := make(map[rune]rune)
	mp[')'] = '('
	mp[']'] = '['
	mp['}'] = '{'

	for _, val := range str {
		switch val {
		case '(', '[', '{':
			stack = append(stack, val)

		case ')', ']', '}':
			if len(stack) == 0 || mp[val] != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}

		}

	}
	fmt.Printf("是否是合格括号%t\n", len(stack) == 0)
	return len(stack) == 0
}

// 查找字符串数组中的最长公共前缀--逐一和第一个元素比较，不断更新第一个元素
func longestCommonPrefix1(str []string) string {

	if len(str) == 0 {
		return ""
	}

	//取第一个元素和所有元素逐一比较
	firstStr := str[0]

	for i := 1; i < len(str); i++ {
		for strings.Index(str[i], firstStr) != 0 {
			firstStr = firstStr[:len(firstStr)-1]
			if firstStr == "" {
				return ""
			}
		}
	}
	fmt.Println("1输出最长公共：%s\n", firstStr)
	return firstStr
}

// 查找字符串数组中的最长公共前缀--同时比较所有元素相同坐标的数据是否相等
func longestCommonPrefix2(str []string) string {

	if len(str) == 0 {
		return ""
	}

	for i := 0; ; i++ {
		for _, val := range str {
			if i >= len(val) || (i < len(val) && str[0][i] != val[i]) {
				fmt.Println("2输出最长公共：%s\n", str[0][:i])
				return str[0][:i]
			}
		}
	}
	return ""
}

//加一 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
