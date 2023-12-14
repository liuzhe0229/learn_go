package main

import "fmt"

type Duck interface {
	walk()
	swimming()
}

type PskDuck struct {
	legs int
}

func (pd *PskDuck) walk() {
	fmt.Println("This psdDuck is walking")
}

func (pd *PskDuck) swimming() {
	fmt.Println("This psdDuck is swimming")
}

func main() {
	var pskDuck Duck = &PskDuck{12}

	pskDuck.walk()
	pskDuck.swimming()
}
