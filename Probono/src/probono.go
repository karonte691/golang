package main

import (
	"os"

	"frasimpatico"
)


func main() {
	file, err := os.Open("test");
	if err != nil {
		panic(err)
	}

	defer file.Close();

	reader, err := frasimpatico.ConvertToUtf8(file)
	if err == nil {
		frasimpatico.SaveFile(reader)
	}
}