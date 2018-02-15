package main

import (
	"fmt"
)

func main() {
	var b byte = 3

	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i *= 3
	i /= 3
	i %= 3

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(y, x)
	x, y = y, x
	fmt.Println(y, x)
}
