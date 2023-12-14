package main

import (
	"fmt"
	"sync"
)

// 实际上就是相当于JS中sync 函数一样，让协程变成同步执行
func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("Execute all go routine")
}
