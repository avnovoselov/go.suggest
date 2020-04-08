package main

import (
	"fmt"
	"math"
)

func forLoop() {
	max := int(math.Pow(2, 3))
	for i := 0; i < max; i++ {
		fmt.Printf("%03b\n", i)
	}
}

func whileLoop() {
	number := 100
	for number > 0 {
		fmt.Println(number)
		number -= 10
	}
}

func main() {
	forLoop()
	whileLoop()
}
