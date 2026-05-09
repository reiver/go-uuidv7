package uuidv7

import (
	"time"
)

// Generate returns a new UUIDv7 for the current time as a [16]byte.
//
// Example usage:
//
//	var uuid [16]byte = uuidv7.Generate()
func Generate() [16]byte {
	var uuid [16]byte
	Put(&uuid)
	return uuid
}

// Put puts a new UUIDv7 for the current time into a [16]byte.
//
// Example usage:
//
//	var uuid [16]byte
//	
//	uuidv7.Put(&uuid)
func Put(uuid *[16]byte) {
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
