package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	// 输出最终结果
	fmt.Printf("计数器最终值: %d\n", atomic.LoadInt64(&counter))
	var input string
	fmt.Scanln(&input)
	fmt.Println("程序结束")
}
