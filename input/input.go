package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Digite se nome")
	nome := bufio.NewReader(os.Stdin)
	text, _ := nome.ReadString('\n')
	fmt.Println(text)
}
