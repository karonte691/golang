package main

import "fmt"

type poligono interface {
	perimetro() int
	area() int
}

type quadrato struct {
	lato int
}

func (q quadrato) perimetro() (int) {
	return q.lato * 4;
}

func (q quadrato) area() (int) {
	return q.lato * q.lato;
}

func m(p poligono) (int) {
	return 1;
}
func main() {
	var q quadrato;

	q.lato = 5;

	//fmt.Println(q.area());

	fmt.Println(q.perimetro());

	m(q)
}




