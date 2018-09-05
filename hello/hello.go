package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/chalchuck/stringutil"
)

func main() {
	fmt.Println("Hello Charles, Welcome to Golang; you are on your path to become a Gopher!!")
	fmt.Println("The time now is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("You have %g Problems", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
