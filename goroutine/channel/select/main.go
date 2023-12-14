package main

import (
	"fmt"
	"time"
)

func g1(c chan struct{}) {
	time.Sleep(time.Second)
	c <- struct{}{}
}

func g2(c chan struct{}) {
	time.Sleep(2 * time.Second)
	c <- struct{}{}
}

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	timer := time.NewTimer(3 * time.Second)

	go g1(c1)
	go g2(c2)

	for {
		select {
		case <-c1:
			fmt.Println("c1 get message")
		case <-c2:
			fmt.Println("c2 get message")
		case <-timer.C:
			fmt.Println("timeout")
			return
		}
	}

}
