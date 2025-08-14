package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}

	list := []int{1, 3, 4, 5, 5, 6, 4, 3, 1}
	singleNum(list)
	singleNum2(list)
	isPalindrome(1234321)
	isPalindrome2(5678998765)
	isValid("{}[](){()}")
	llist := []string{"234", "234", "2346456", "23"}
	longestCommonPrefix1(llist)
	longestCommonPrefix2(llist)
	plusOne([]int{9, 9, 9, 9, 9})
	twoSum([]int{2, 7, 11, 15}, 18)
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	mergNum := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(mergNum)

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

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

// 加一 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(num []int) []int {

	for i := len(num) - 1; i >= 0; i-- {
		num[i]++
		if num[i] < 10 {
			return num
		}
		num[i] = 0
	}
	result := append([]int{1}, num...)
	fmt.Println("加一最终结果\n", result)
	return result
}

// twoSum 找出数组中和为目标值的两个整数的下标
func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int)
	for index, val := range nums {

		needNum := target - val
		if i, exist := numMap[needNum]; exist {
			fmt.Println("返回两数坐标:", []int{i, index})
			return []int{i, index}
		}
		numMap[val] = index
	}
	return nil
}

// 删除有序数组中的重复项，返回唯一元素的个数k
func removeDuplicates(nums []int) int {
	// 处理空数组情况
	if len(nums) == 0 {
		return 0
	}

	// 慢指针，指向当前最后一个唯一元素的位置
	slow := 0

	// 快指针，遍历整个数组
	for fast := 1; fast < len(nums); fast++ {
		// 当快指针发现与慢指针指向的元素不同时
		if nums[fast] != nums[slow] {
			// 慢指针向前移动一位
			slow++
			// 将快指针指向的元素复制到慢指针位置
			fmt.Print(nums[slow], nums[fast])
			nums[slow] = nums[fast]
			fmt.Print("去重复，交换后", nums[slow], nums[fast])
			fmt.Println("\n")

		}
		// 如果元素相同，则快指针继续前进，慢指针不动

	}

	// 唯一元素的个数是慢指针索引 + 1
	return slow + 1
}

// 合并区间
func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return intervals
	}
	//从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {

		last := result[len(result)-1]
		if last[1] >= intervals[i][0] {
			last[1] = max(intervals[0][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}

	fmt.Printf("最终%d", result)
	return result
}

// 辅助函数：取两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
