package main

import (
	"fmt"
	"strconv"
)

func main() {
	const PI float32 = 3.1415926

	const (
		UNKNOWN = 1
		FEMALE
		MALE
		animal
	)

	const (
		ERR1 = iota
		ERR2 = iota + 4
		ERR3
		ERR4 = 5
		ERR5 = iota
	)

	fmt.Println(MALE)

	var myInt = 12
	fmt.Println(strconv.Itoa(myInt))

	// 单引号就是byte类型
	var c byte
	c = 'a'
	fmt.Println(c)
	fmt.Printf("c=%c \n", c)

	var age = 12
	var name = "liu Zhe"
	var address = "changsha"

	fmt.Printf("%s今年%d住在%s \n", name, age, address)

	var message = fmt.Sprintf("%s今年%d住在%s", name, age, address)
	fmt.Println(message)

	var m = map[string]string{
		"go": "Go language",
	}
	fmt.Println(m["go"])
}
