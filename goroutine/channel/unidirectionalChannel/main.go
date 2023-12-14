package main

import (
	"fmt"
	"time"
)

// 单向channel目前发现的用法其实就是在函数里面用来限制channel只有一个功能
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	time.Sleep(time.Second)
}
