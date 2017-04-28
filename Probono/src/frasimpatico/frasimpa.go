package frasimpatico

import (
	iconv "github.com/djimenez/iconv-go"
	"os"
	"io"
)


func ConvertToUtf8(f *os.File) (io.Reader, error){

	reader, err := iconv.NewReader(f, "big5", "utf-8")

	if err != nil {
		return nil, err
	}

	return reader, nil
}

func SaveFile(reader io.Reader) (error) {
	fo, err := os.Create("test_utf8")
	if err != nil {
		return err
	}
	defer fo.Close()
	_, errCopy := io.Copy(fo, reader)

	return errCopy
}
