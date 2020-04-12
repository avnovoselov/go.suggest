package loop

import (
	"fmt"
	"math"
	"runtime"
	"strings"
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

func newtonSqrt(number float64, accuracy int) float64 {
	assumption := number / 3
	for accuracy > 0 {
		assumption = assumption - (math.Pow(assumption, 2)-number)/2/assumption
		accuracy--
	}

	return assumption
}

func whatImThinkAboutYourOS(os string) string {
	var resume string
	switch {
	case strings.Contains(os, "windows"):
		resume = "normal"
	case strings.Contains(os, "mac"):
		resume = "good"
	case strings.Contains(os, "linux"):
		resume = "nice"
	default:
		resume = "wtf?"
	}

	return resume
}

func Loop() {
	forLoop()
	whileLoop()
	fmt.Println(newtonSqrt(36, 100))
	fmt.Println("Your OS is:", whatImThinkAboutYourOS(runtime.GOOS))
}
