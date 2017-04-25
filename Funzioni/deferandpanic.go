package main

import (
	"fmt"
	"errors"
)

func main() {
	i := 5;

	resdue, _ := ritornaPerCinque(i);

	fmt.Println(resdue)

	esegui()

	fmt.Println("Andiamo avanti")
}

func esegui() {
	defer func() {
		r := recover();

		if r != nil {
			fmt.Println("ritona per 5 ha panicato, ma niente panico")
		}
	}();

	_, err := ritornaPerCinque(0);

	if err != nil {
		panic(err)
	}


}

func ritornaPerCinque(i int) (int, error) {
	if i == 0 {
		return -1, errors.New("Il numero non può essere 0!")
	}

	return i * 3, nil;
}