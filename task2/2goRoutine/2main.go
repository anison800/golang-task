package main

import (
	"fmt"
	"sync"
	"time"
)

// 执行单个任务并返回执行时间
func runTask(id int, task func()) time.Duration {
	start := time.Now()
	defer func() {
		// 捕获任务中的错误
		if err := recover(); err != nil {
			fmt.Printf("任务 %d 执行出错: %v\n", id, err)
		}
	}()

	task() // 执行任务
	return time.Since(start)
}

// 调度器：并发执行所有任务并统计时间
func scheduleTasks(tasks ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	// 并发执行每个任务
	for i, task := range tasks {
		go func(id int, t func()) {
			defer wg.Done()
			duration := runTask(id, t)
			fmt.Printf("任务 %d 执行完成，耗时: %v\n", id, duration)
		}(i+1, task) // 传递当前索引作为任务ID
	}

	wg.Wait() // 等待所有任务完成
}

func main() {
	// 示例任务
	task1 := func() { time.Sleep(200 * time.Millisecond) }
	task2 := func() { time.Sleep(300 * time.Millisecond) }
	task3 := func() { time.Sleep(150 * time.Millisecond) }
	task4 := func() {
		// 模拟错误任务
		time.Sleep(100 * time.Millisecond)
		panic("任务内部错误")
	}

	// 执行所有任务
	fmt.Println("开始执行任务...")
	start := time.Now()
	scheduleTasks(task1, task2, task3, task4)

	fmt.Printf("\n所有任务执行完毕，总耗时: %v\n", time.Since(start))
}
