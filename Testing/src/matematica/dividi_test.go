package matematica_test

import "testing"
import "matematica"

func TestDivisibilePerDue(t *testing.T) {
	v := 2;
	x := 3;

	var res bool;

	res = matematica.DivisibilePerDue(v);

	if res == false {
		t.Error("Doveva tornare true, invece è tornato ", res)
	}

	res = matematica.DivisibilePerDue(x);

	if res != false {
		t.Error("Doveva tornare false, invece è tornato ", res)
	}
}
