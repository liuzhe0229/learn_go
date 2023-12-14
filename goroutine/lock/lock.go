package main

import (
	"fmt"
	"sync"
)

// 使用lock 来保证 读值，运算，赋值是原子操作
var wg sync.WaitGroup
var lock sync.Mutex
var total = 0

func add() {
	defer wg.Done()
	for i := 0; i <= 10000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i <= 10000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}

func main() {

	wg.Add(2)
	go add()
	go sub()
	wg.Wait()

	fmt.Println(total)
}
