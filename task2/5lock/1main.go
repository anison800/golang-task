package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
	counter := &counter{}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.increment()
			}
		}()
	}
	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Printf("最终计数器值: %d\n", counter.getCount())
}

type counter struct {
	count int
	mu    sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
