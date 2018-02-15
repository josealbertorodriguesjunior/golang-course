package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if sum%2 == 0 {
			fmt.Println("Par ")
			fmt.Println(sum)
		} else {
			fmt.Println("Impar ")
			fmt.Println(sum)
		}
	}
}
