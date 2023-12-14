package main

import (
	"fmt"
	"time"
)

// 使用无缓冲channel 时必须要要在 go routine中消费
// 不然就会死锁
// happen-before 机制能保证channel 的传值一定发生在取值之前
func main() {
	msg := make(chan int, 2)

	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
	}(msg)

	msg <- 1
	msg <- 2
	msg <- 3

	close(msg) // 一个chanel 被关闭之后就不能再写值进去了 会报panic

	time.Sleep(3 * time.Second)
}
