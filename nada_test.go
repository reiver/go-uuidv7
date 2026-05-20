package uuidv7_test

import (
	"testing"

	"github.com/reiver/go-uuidv7"
)

func TestNada(t *testing.T) {
	const expected string = "00000000-0000-0000-0000-000000000000"

	actual := uuidv7.String(uuidv7.Nada())

	if expected != actual {
		t.Errorf("The actual UUIDv7 string for 'min' is not what was expected.")
		t.Logf("EXPECTED: %s", expected)
		t.Logf("ACTUAL:   %s", actual)
		return
	}
}
