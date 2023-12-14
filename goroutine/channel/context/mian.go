package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context, funcName string) {
	defer wg.Done()
	if ctx.Value("traceId") != nil {
		fmt.Println(ctx.Value("traceId"))
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("***%s***Stop the cpu watch \n", funcName)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("***%s***Watching the Cpu\n", funcName)
		}
	}
}

// context 有三种方法 withCancel, withValue, withTimeout
/*
withCancel 的cancel方法可以取消子ctx
如果调用cancel()，那么ctx1.done()channel就会被传值
*/
/*
withTimeout 可以自带cancel 通过指定的时间自动给 ctx.done()传值
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx1, _  := context.WithCancel(ctx)
	wg.Add(2)

	go cpuInfo(ctx, "withCancel")

	ctxTimerOut, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go cpuInfo(ctxTimerOut, "withTimeout")

	ctxWithValue := context.WithValue(ctxTimerOut, "traceId", "12345678")
	go cpuInfo(ctxWithValue, "ctxWithValue")

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}
