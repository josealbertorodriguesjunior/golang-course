package main

import (
	"fmt"
)

func main() {
	i := 1

	//Go não suporta aritimética com ponteiros

	var p *int = nil
	p = &i
	*p++
	fmt.Println(p, *p, i, &i)
}
