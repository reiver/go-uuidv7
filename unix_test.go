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
			Expected: 0xFFFFFFFF_FFFF, // == 281474976710655
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
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(9998),
			Expected:                              9998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(9999),
			Expected:                              9999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(10000),
			Expected:                              10000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(10001),
			Expected:                              10001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(99998),
			Expected:                              99998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(99999),
			Expected:                              99999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(100000),
			Expected:                              100000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(100001),
			Expected:                              100001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(999998),
			Expected:                              999998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(999999),
			Expected:                              999999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1000000),
			Expected:                              1000000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1000001),
			Expected:                              1000001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(9999998),
			Expected:                              9999998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(9999999),
			Expected:                              9999999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(10000000),
			Expected:                              10000000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(10000001),
			Expected:                              10000001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(99999998),
			Expected:                              99999998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(99999999),
			Expected:                              99999999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(100000000),
			Expected:                              100000000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(100000001),
			Expected:                              100000001,
		},



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(999999998),
			Expected:                              999999998,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(999999999),
			Expected:                              999999999,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1000000000),
			Expected:                              1000000000,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(1000000001),
			Expected:                              1000000001,
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



		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(0xFFFFFFFF_FFFF),
			Expected:                              0xFFFFFFFF_FFFF,
		},
		{
			UUIDv7: uuidv7.GenerateAtUnixTimeMilli(281474976710655),
			Expected:                              281474976710655,
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
