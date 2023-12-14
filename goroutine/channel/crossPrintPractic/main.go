package main

import (
	"fmt"
	"time"
)

func printNumber(numberChan, letterChan chan bool) {
	i := 1

	for {
		<-numberChan
		fmt.Printf("%d%d", i, i+1)
		i = i + 2
		letterChan <- true
	}
}

func printLetter(numberChan, letterChan chan bool) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(str); i += 2 {
		<-letterChan
		fmt.Print(str[i : i+2])
		numberChan <- true
	}
}

func main() {
	numberChan := make(chan bool)
	letterChan := make(chan bool)

	go printLetter(numberChan, letterChan)
	go printNumber(numberChan, letterChan)

	numberChan <- true

	time.Sleep(5 * time.Second)
}
