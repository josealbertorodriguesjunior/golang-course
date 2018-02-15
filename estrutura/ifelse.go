package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com", nota)
	} else {
		fmt.Println("Reprovado com", nota)
	}
}

func main() {
	imprimirResultado(5.0)
}
