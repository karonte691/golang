package main

import "fmt"

func main() {
	var arraySemplice  = [5]int{1,2,3,4,5};

	fmt.Println("Array semplice: ", arraySemplice)

	var slice = make([]int, 5);

	slice[0] = 1;
	slice[1] = 2;
	slice[2] = 3;
	slice[3] = 4;
	slice[4] = 5;

	fmt.Println("Slice: ", slice)

	//append qualcosa
	slice = append(slice, "6");

	fmt.Println("Slice dopo append: ", slice);

	//copy
	slicecopy := make([]int, len(slice));

	copy(slicecopy, slice)

	fmt.Println("Slice copiato: ", slicecopy);



}