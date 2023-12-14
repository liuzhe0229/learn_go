package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float32
}

func main() {
	var name string
	name = "zhe.liu"
	age := 1

	var (
		address  string
		birthday = 12312312312
	)

	address = "changsha"
	fmt.Println(name, age)
	fmt.Println(address, birthday)

	var slice1 []string
	var arr = [5]string{"changsha", "yueyang"}
	slice1 = arr[0:2]

	fmt.Println(slice1)

	var person1 Person = Person{"Micheal", 18, 182.0}

	fmt.Println(person1)
	fmt.Printf("%+v \n", person1)
}
