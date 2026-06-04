package uuidv7

var min [16]byte = [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0x80,0x00,0x00,0x00,0x00,0x00,0x00,0x00}

// Min returns the minimum value for a UUIDv7.
// I.e., 00000000-0000-7000-8000-000000000000.
//
// Note that not all values between the minimum UUIDv7 and the maximum UUIDv7 are valid UUIDv7.
// For example, 11111111-1111-1111-1111-111111111111 is between the minimum UUIDv7 and the maximum UUIDv7 but is NOT a valid UUIDv7.
// You can use [Is] and [IsString] to determine if something is a valid UUIDv7.
func Min() [16]byte {
	return min
}
