package uuidv7

import (
	"time"

	"codeberg.org/reiver/go-erorr"
)

const uuidStringLength = 36

// Generate returns a new UUIDv7 for the current time as a [16]byte.
//
// Example usage:
//
//	var uuid [16]byte = uuidv7.Generate()
func Generate() [16]byte {
	var uuid [16]byte
	GeneratePut(&uuid)
	return uuid
}

// GeneratePut puts a new UUIDv7 for the current time into a [16]byte.
//
// Example usage:
//
//	var uuid [16]byte
//	
//	uuidv7.GeneratePut(&uuid)
func GeneratePut(uuid *[16]byte) {
	if nil == uuid {
		return
	}

	{
		now := time.Now().UnixMilli()

		uuid[0] = byte(now >> 40)
		uuid[1] = byte(now >> 32)
		uuid[2] = byte(now >> 24)
		uuid[3] = byte(now >> 16)
		uuid[4] = byte(now >>  8)
		uuid[5] = byte(now      )
	}

	{
		rnd := randomUint64()

		uuid[ 6] = byte(rnd >> 56)
		uuid[ 7] = byte(rnd >> 48)
		uuid[ 8] = byte(rnd >> 40)
		uuid[ 9] = byte(rnd >> 32)
		uuid[10] = byte(rnd >> 24)
		uuid[11] = byte(rnd >> 16)
		uuid[12] = byte(rnd >>  8)
		uuid[13] = byte(rnd      )
	}

	{
		rnd := randomUint64()

		uuid[14] = byte(rnd >> 8)
		uuid[15] = byte(rnd     )
	}

	{
		uuid[6] = (uuid[6] & 0x0F) | 0x70
		uuid[8] = (uuid[8] & 0x3F) | 0x80
	}
}

func Is(uuid [16]byte) bool {
	var m byte = (uuid[6] & 0xF0)

	if 0x70 != m {
		return false
	}

	var n byte = (uuid[8] & 0xC0)

	if 0x80 != n && 0x90 != n && 0xa0 != n && 0xb0 != n {
		return false
	}

	return true
}

func IsString(str string) bool {
	err := ValidateString(str)
	return nil == err
}

const (
	ErrHyphenMissing   = erorr.Error("hypen missing")
	ErrLengthWrong     = erorr.Error("length wrong")
	ErrNotUUID         = erorr.Error("not UUID")
	ErrNotVersion7UUID = erorr.Error("not version 7 UUID")
)

func ValidateString(str string) error {
	if uuidStringLength != len(str) {
		return ErrLengthWrong
	}

	if '-' != str[8] {
		return ErrHyphenMissing
	}
	if '-' != str[13] {
		return ErrHyphenMissing
	}
	if '-' != str[18] {
		return ErrHyphenMissing
	}
	if '-' != str[23] {
		return ErrHyphenMissing
	}

	for i := range uuidStringLength {
		if 8 == i || 13 == i || 18 == i || 23 == i {
			continue
		}

		if str[i] < '0' || ('9' < str[i] && str[i] < 'A') || ('F' < str[i] && str[i] < 'a') || 'f' < str[i] {
			return ErrNotUUID
		}
	}

	if '7' != str[14] {
		return ErrNotVersion7UUID
	}
	if '8' != str[19] && '9' != str[19] && 'A' != str[19] && 'B' != str[19] && 'a' != str[19] && 'b' != str[19] { // b10xx
		return ErrNotVersion7UUID
	}

	return nil
}
