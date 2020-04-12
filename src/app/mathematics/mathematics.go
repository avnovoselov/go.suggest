package mathematics

import (
	"fmt"
	"math/rand"
)

func randomInteger() int {
	rand.Seed(123927358235)

	return rand.Intn(100)
}

func sumAndDifference(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b

	return
}

func deferred(a, b int) {
	defer func(a, b int) {
		// keep it to storage
		fmt.Println("Store:", a+b)
	}(a, b)

	fmt.Println("Return:", a-b)
}

func Mathematics() {
	fmt.Println("Random:", randomInteger())

	sum, difference := sumAndDifference(10, 2)
	fmt.Println("Sum:", sum, "difference:", difference)

	deferred(3, 1)
}
