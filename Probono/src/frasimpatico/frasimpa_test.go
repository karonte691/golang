package frasimpatico_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"unicode/utf8"
	"frasimpatico"
	"os"
)

func TestConvertToUtf8(t *testing.T) {
	asserzione := assert.New(t)

	file, err := os.Open("testFiles/frasimpa_convertutf8.big5.txt");
	if err != nil {
		asserzione.Error(err)
	}
	defer file.Close();

	reader, err := frasimpatico.ConvertToUtf8(file)

	byteString := []byte{}
	buf := make([]byte, 40)

	for n := 1; n > 0; {
		n, err = reader.Read(buf)
		if err != nil {
			break
		}
		byteString = append(byteString, buf[:n]...)
	}

	s := string(byteString[:])
	asserzione.NotEmpty(s)
	asserzione.True(utf8.ValidString(s))

	t.Log(utf8.ValidString(s), s)
}