package main

import "fmt"

func math_t(a int, b int) (int, int, int, int) {
	sum := a + b
diff:
	+a - b
	mult := a * b
	div := a / b
	return sum, diff, mult, div
}

func main() {
	sum, diff, mult, div := math_t(6, 8)
	fmt.Println(sum, diff, mult, div)
}
