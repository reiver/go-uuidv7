package uuidv7

func MustParse(str string) [16]byte {
	uuid, err := Parse(str)
	if nil != err {
		panic(err)
	}

	return uuid
}

func Parse(str string) ([16]byte, error) {
	var uuid [16]byte
	err := ParsePut(&uuid, str)
	return uuid, err
}

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
