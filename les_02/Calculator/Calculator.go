package main

import (
	"fmt"
)

func Cal(a string, b, c int) float32 {
	switch a {
	case "+":
		return float32(b + c)
	case "-":
		return float32(b - c)
	case "*":
		return float32(b * c)
	case "/":
		if c == 0 {
			return 0
		}
		return float32(b / c)
	}
	return 0
}

func main() {
	fmt.Println(Cal("+", 10, 4))
}
