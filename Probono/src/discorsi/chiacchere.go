package discorsi

import "errors"

func SalutiBase(saluto []string, risposta[] string) ([]string, error) {
	if len(saluto) == 0 || len(risposta) == 0 {
		return nil, errors.New("Non si può fare un discorso da soli")
	}

	var res = append(saluto,risposta...)

	return res, nil
}