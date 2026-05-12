package uuidv7

// UnixTimeMilli returns the unix-time timestamp, in milliseconds, for a UUIDv7.
//
// With UUIDv7, the first 48-bits / 6-bytes is a unix-time timestamp in milliseconds.
//
// Example usage:
//
//	var uuid [16]byte = uuidv7.MustParse("019e1a5e-b6d4-757c-8890-9d2969f29d7e")
//	
//	unixTimeStamp := uuidv7.UnixTimeMilli(uuid)
func UnixTimeMilli(uuid [16]byte) (int64, error) {
	if !Is(uuid) {
		return 0, ErrNotVersion7UUID
	}

	return int64(
		(uint64(uuid[0]) << 40) |
		(uint64(uuid[1]) << 32) |
		(uint64(uuid[2]) << 24) |
		(uint64(uuid[3]) << 16) |
		(uint64(uuid[4]) << 8)  |
		(uint64(uuid[5])     ),
	), nil
}
