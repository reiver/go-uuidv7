package uuidv7_test

import (
	"testing"

	"github.com/reiver/go-uuidv7"
)

func TestMax(t *testing.T) {
	const expected string = "ffffffff-ffff-7fff-bfff-ffffffffffff"

	actual := uuidv7.String(uuidv7.Max())

	if expected != actual {
		t.Errorf("The actual UUIDv7 string for 'max' is not what was expected.")
		t.Logf("EXPECTED: %s", expected)
		t.Logf("ACTUAL:   %s", actual)
		return
	}
}
