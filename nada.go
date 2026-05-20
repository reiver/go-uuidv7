package uuidv7

// Nada returns the UUID 00000000-0000-0000-0000-000000000000.
// Note that 00000000-0000-0000-0000-000000000000 is NOT a valid UUIDv7.
func Nada() [16]byte {
	return [16]byte{}
}
