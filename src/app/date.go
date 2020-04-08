package app

import (
	"fmt"
	"time"
)

func PrintNow() {
	fmt.Println(fmt.Sprintf("Сейчас: %[1]s", time.Now().Format("02-01-2016 03:05:04")))
}
