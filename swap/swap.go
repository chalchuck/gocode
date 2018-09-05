package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "Gopher")
	fmt.Println(a, b)

	fmt.Println("Hello"[1])
	fmt.Println(len("Hello world"))
}
