package date

import (
	"fmt"
	"time"
)

func Date() {
	fmt.Println(fmt.Sprintf("Сейчас: %[1]s", time.Now().Format("2006-01-02 03:04:05")))
}
