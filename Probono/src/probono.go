package main

import (
	"discorsi"
	"fmt"
)


func main() {
	var s  = []string{"ciao","come","va", "?"};
	var t  = []string{"bene", "grazie", "te"}

	var chiaccherata, _ = discorsi.SalutiBase(s,t);

	fmt.Println("Un semplice discorso la mattina: ", chiaccherata)
}