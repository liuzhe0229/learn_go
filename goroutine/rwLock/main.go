package main

import (
	"fmt"
	"sync"
	"time"
)

// 一开始程序正常运行
// 进入写协程时 先sleep 5s,此时读协程依然在正常运行 不停打印
// 写协程加锁之后，读协程就停止运行
// 写协程接触读写锁之后，读协程继续执行
func main() {
	var rwLock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		rwLock.Lock()
		fmt.Println("------Get writ lock-------")
		time.Sleep(5 * time.Second)
		rwLock.Unlock()
	}()

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				rwLock.RLock()
				fmt.Printf("Number %d excute ready process \n", i)
				time.Sleep(500 * time.Millisecond)
				rwLock.RUnlock()
			}
		}(i)
	}
	wg.Wait()

}
