package main

import (
	. "example.com/suggest/src/app/suggest"
	. "example.com/suggest/src/app/suggest/environment"
	"log"
)

func main() {
	log.Println("Start suggest service")

	loadError, variables := LoadVariables()

	if loadError == nil {
		suggest := Suggest{}

		suggest.Create(variables).Run()
	}
}
