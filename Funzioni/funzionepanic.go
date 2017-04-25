package main

import (
	"fmt"
	"errors"
)

func main() {
	i := 5;

	_, err := ritornaPerQuattro(0);

	if err != nil {
		panic(err)
	}

	resdue, _ := ritornaPerQuattro(i);

	fmt.Println(resdue)
}

func ritornaPerQuattro(i int) (int, error) {
	if i == 0 {
		return -1, errors.New("Il numero non può essere 0!")
	}

	return i * 3, nil;
}