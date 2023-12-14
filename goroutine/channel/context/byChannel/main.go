package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(stop chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("Stop the cpu watch")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Watching the Cpu")
		}
	}
}

func main() {
	stop := make(chan struct{})

	wg.Add(1)

	go cpuInfo(stop)
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	wg.Wait()
}
