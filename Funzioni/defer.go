package main

import (
	"fmt"
)

func main() {

	fmt.Println("frase 1")
	defer scriviFraseTre()
	fmt.Println("Frase 2")

}

func scriviFraseTre() {
	fmt.Println("Frase 3")
}