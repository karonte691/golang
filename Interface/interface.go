package main

import "fmt"

type quadratoFormule interface {
	perimetro() int
	area() int
}

type quadrato struct {
	lato int
}

func main() {

	var q quadrato;

	q.lato = 5;

	fmt.Println(q.area());
	fmt.Println(q.perimetro());
}



func (q quadrato) perimetro() (int) {
	return q.lato * 4;
}

func (q quadrato) area() (int) {
	return q.lato * q.lato;
}
