package matematica

import "errors"

func Moltiplica(num int, perQuanto int) (int, error) {
	if perQuanto == 0 {
		//tanto per testing
		return -1, errors.New("Valore per quanto a 0!")
	}

	if num == 0 {
		//sempre tanto per testing
		return -1, errors.New("Valore iniziale uguale a 0")
	}

	return num * perQuanto, nil
}