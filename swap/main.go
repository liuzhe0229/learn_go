package main

import "fmt"

func swap(a, b *int) {
	a, b = b, a
	fmt.Println(*a, *b)
	t := *a
	*a = *b
	*b = t
}

func main() {
	var a, b = 1, 2

	swap(&a, &b)

	fmt.Println(a, b)
}
