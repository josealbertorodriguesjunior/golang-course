package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("Um byte é ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("Tipo de i1", reflect.TypeOf(i1))

	var i2 rune = 'a'

	fmt.Println("i2 é do tipo", reflect.TypeOf(i2))
	fmt.Println(i2)

	fmt.Println(len("Isso pode ser uma coisa muito louca"))
}
