package discorsi

import "errors"

func Discorso(t []string, s []string) ([]string, error) {

	lunghezzaDiscorso := len(t) + len(s)

	discorso := make([]string, 0)

	if lunghezzaDiscorso == 0 {
		return discorso, errors.New("Il discorso non può essere vuoto")
	}

	for j := 0; j < len(t); j++ {
		discorso = append(discorso, t[j])
	}

	for j := 0; j < len(s); j++ {
		discorso = append(discorso, s[j])
	}

	return discorso, nil;

}

