package uuidv7_test

import (
	"testing"

	"github.com/reiver/go-uuidv7"
)

func TestUnixTimeMilli(t *testing.T) {
	tests := []struct{
		UUIDv7 [16]byte
		Expected int64
	}{
		{
			UUIDv7:   uuidv7.Min(),
			Expected: 0,
		},
		{
			UUIDv7:   uuidv7.Max(),
			Expected: 281474976710655,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(0),
			Expected:                              0,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1),
			Expected:                              1,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(2),
			Expected:                              2,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(3),
			Expected:                              3,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(4),
			Expected:                              4,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(5),
			Expected:                              5,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(98),
			Expected:                              98,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(99),
			Expected:                              99,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(100),
			Expected:                              100,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(101),
			Expected:                              101,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(998),
			Expected:                              998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(999),
			Expected:                              999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1000),
			Expected:                              1000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1001),
			Expected:                              1001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1778302347),
			Expected:                              1778302347,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1780531816),
			Expected:                              1780531816,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(98765432101234),
			Expected:                              98765432101234,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(222222222222000),
			Expected:                              222222222222000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTime(222222222222),
			Expected:                         222222222222000,
		},
	}

	for testNumber, test := range tests {
		actual, err := uuidv7.UnixTimeMilli(test.UUIDv7)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("UUIDv7: %s", uuidv7.String(test.UUIDv7))
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual millisecond-resolution unix-time it not as expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("UUIDv7: %s", uuidv7.String(test.UUIDv7))
			continue
		}
	}
}
