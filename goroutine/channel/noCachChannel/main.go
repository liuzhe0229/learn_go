package main

import (
	"fmt"
	"time"
)

// 使用无缓冲channel 时必须要要在 go routine中消费
// 不然就会死锁
// happen-before 机制能保证channel 的传值一定发生在取值之前
func main() {
	msg := make(chan string, 0)

	go func(msg chan string) {
		data := <-msg
		fmt.Println(data)
	}(msg)

	go func(msg chan string) {
		time.Sleep(3 * time.Second)
		msg <- "zhe.liu"
	}(msg)

	time.Sleep(10 * time.Second)
}
