package main

import "fmt"

func main() {
	var input string;
	//var s  = 5;


	fmt.Println("Per favore, inserisci una parola...: ");

	fmt.Scanln(&input)

	for j := 1; j < 10; j++ {
		fmt.Println("La tua parola " + input + "ciclata con un for normale")
	}

	i := 0;

	for i <= 3 {
		fmt.Println("La tua parola " + input + "ciclata con un for breve")
		i++;
	}


}