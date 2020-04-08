package app

import "math/rand"

func Sum(x, y int) int {
	return x + y
}

func RandomInteger() int {
	rand.Seed(123927358235)

	return rand.Intn(100)
}

func SumAndDifference(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b

	return
}
