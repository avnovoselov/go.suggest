package app

import "fmt"

const (
	C          bool = false
	Java       bool = false
	PHP        bool = true
	JavaScript bool = true
)

func YouNow() {
	//noinspection GoBoolExpressions
	fmt.Println(fmt.Sprintf("C: %[1]t, Java: %[2]t, PHP: %[3]t, JavaScript: %[4]t", C, Java, PHP, JavaScript))
}
