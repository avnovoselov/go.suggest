package text

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func greeting(Name string) {
	fmt.Printf("Hello, %s\n", Name)
}

func Text() {
	greeting("Golang")
	greeting, name := swap("Hello", "World")
	fmt.Println(greeting, name)
}
