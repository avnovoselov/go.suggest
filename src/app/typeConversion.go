package app

import (
	"fmt"
	"strconv"
)

func Run() {
	var (
		someString                       = "Foo"
		someStringContainsNumber         = "123"
		number                   int32   = 1024
		single                   float32 = 2.57
	)

	convertedStringToInt, _ := strconv.Atoi(someString)
	fmt.Println(fmt.Sprintf("Bad string to int32: %[1]d", convertedStringToInt))

	convertedStringToInt, _ = strconv.Atoi(someStringContainsNumber)
	fmt.Println(fmt.Sprintf("Good string to int32: %[1]d", convertedStringToInt))

	fmt.Println(fmt.Sprintf("Int32 to float32: %[1]f", float32(number)))
	fmt.Println(fmt.Sprintf("Float32 to int32: %[1]d", int32(single)))
}
