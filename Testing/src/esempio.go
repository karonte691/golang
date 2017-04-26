package main

import "fmt"
import "matematica"

func main() {
	fmt.Println("Inserisci un numero")

	var input int;

	fmt.Scanf("%d", &input)

	if matematica.DivisibilePerDue(input) {
		fmt.Println("Il numero inserito è divisibile per due!")
	} else {
		fmt.Println("Il numero inserito NON è divisibile per due")
	}

}

