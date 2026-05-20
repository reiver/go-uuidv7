package uuidv7_test

import (
	"testing"

	"github.com/reiver/go-uuidv7"
)

func TestMin(t *testing.T) {
	const expected string = "00000000-0000-7000-8000-000000000000"

	actual := uuidv7.String(uuidv7.Min())

	if expected != actual {
		t.Errorf("The actual UUIDv7 string for 'min' is not what was expected.")
		t.Logf("EXPECTED: %s", expected)
		t.Logf("ACTUAL:   %s", actual)
		return
	}
}
