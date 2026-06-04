package uuidv7

var max [16]byte = [16]byte{0xff,0xff,0xff,0xff,0xff,0xff,0x7f,0xff,0xbf,0xff,0xff,0xff,0xff,0xff,0xff,0xff}

// Max returns the maximum value for a UUIDv7.
// I.e., ffffffff-ffff-7fff-bfff-ffffffffffff.
//
//
// Note that not all values between the minimum UUIDv7 and the maximum UUIDv7 are valid UUIDv7.
// For example, 11111111-1111-1111-1111-111111111111 is between the minimum UUIDv7 and the maximum UUIDv7 but is NOT a valid UUIDv7.
// You can use [Is] and [IsString] to determine if something is a valid UUIDv7.
func Max() [16]byte {
	return max
}
