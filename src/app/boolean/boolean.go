package boolean

import "fmt"

const (
	C          bool = false
	Java       bool = false
	PHP        bool = true
	JavaScript bool = true
)

func Boolean() {
	//noinspection GoBoolExpressions
	fmt.Printf("C: %[1]t, Java: %[2]t, PHP: %[3]t, JavaScript: %[4]t\n", C, Java, PHP, JavaScript)
}
