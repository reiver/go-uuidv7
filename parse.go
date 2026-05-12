package uuidv7

// MustParse is similar to [Parse] expect it panics if there is an error.
//
// Example usage:
//
//	uuid := uuidv7.MustParse("ed7ba470-8e54-465e-825c-99712043e01c")
//
// See also:
//
//	• [Parse]
//	• [ParsePut]
func MustParse(str string) [16]byte {
	uuid, err := Parse(str)
	if nil != err {
		panic(err)
	}

	return uuid
}

// Parse parses a string for a UUIDv7 and returns a UUIDv7 as a [16]byte if the string contained a valid UUIDv7, else it returns an error.
//
// Example usage:
//
//	uuid, err := uuidv7.Parse("ed7ba470-8e54-465e-825c-99712043e01c")
//
// See also:
//
//	• [MustParse]
//	• [ParsePut]
func Parse(str string) ([16]byte, error) {
	var uuid [16]byte
	err := ParsePut(&uuid, str)
	return uuid, err
}

// ParsePut parses a string for a UUIDv7 and puts the UUIDv7 into a [16]byte if the string contained a valid UUIDv7, else it returns an error.
//
// Example usage:
//
//	var uuid [16]byte
//	err := uuidv7.ParsePut(&uuid, "ed7ba470-8e54-465e-825c-99712043e01c")
//
// See also:
//
//	• [MustParse]
//	• [Parse]
func ParsePut(uuid *[16]byte, str string) error {
	err := ValidateString(str)
	if nil != err {
		return err
	}

	var index int
	for i:=0; i < len(str); i += 2 {
		if '-' == str[i] {
			i++
		}

		high := decodeHex(str[i])
		low := decodeHex(str[1+i])

		uuid[index] = (high << 4) | low
		index++
	}

	return nil
}

func decodeHex(b byte) byte {
	switch {
	case '0' <= b && b <= '9':
		return b - '0'

	case 'a' <= b && b <= 'f':
		return b - 'a' + 10

	case 'A' <= b && b <= 'F':
		return b - 'A' + 10

	default:
		return 0
	}
}
