package discorsi_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"discorsi"
)

func TestDiscorso(t *testing.T) {

	var discorsoAtteso = make([]string, 0);

	var sS = []string{"ciao", "come", "va", "?"};
	var tS = []string{"bene", "grazie", "te"}

	discorsoAtteso = append(discorsoAtteso, sS...)
	discorsoAtteso = append(discorsoAtteso, tS...)

	var asserzione = assert.New(t)

	discorsoCompleto, _ := discorsi.Discorso(sS, tS);

	asserzione.Equal(discorsoAtteso, discorsoCompleto)

}


func TestDiscorsoErrore(t *testing.T) {

	var discorsoAtteso = make([]string, 0)

	var asserzione = assert.New(t)

	var s []string
	var c []string

	discorsoCompleto, err := discorsi.Discorso(s, c)

	asserzione.Equal(discorsoAtteso, discorsoCompleto)

	asserzione.NotNil(err)
	asserzione.Equal("Il discorso non può essere vuoto", err.Error())

}


