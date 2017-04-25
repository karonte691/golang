package main

import (
	"fmt"
	"errors"
)

func main() {
	i := 5;

	_, err := ritornaPerTre(0);

	if err != nil {
		fmt.Println(err)
	}

	resdue, _ := ritornaPerTre(i);

	fmt.Println(resdue)
}

func ritornaPerTre(i int) (int, error) {
	if i == 0 {
		return -1, errors.New("Il numero non può essere 0!")
	}

	return i * 3, nil;
}
