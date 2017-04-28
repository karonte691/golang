package matematica_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"matematica"
)


func TestMoltiplica(t *testing.T) {
	asserzione := assert.New(t)

	x := 3;
	j := 2;

	var res, _ = matematica.Moltiplica(x, j);
	asserzione.Equal(5, res, "Dovrebbe essere 6")

}

func TestMoltiplicaWillFailForPerQuantoSetToZero(t *testing.T) {
	asserzione := assert.New(t)

	x := 3;
	j := 0;

	var res, err = matematica.Moltiplica(x, j);

	asserzione.Equal(-1, res)
	asserzione.NotNil(err)

	asserzione.Equal("Valore per quanto a 0!", err.Error(), "Per quanto deve andare in errore")

}