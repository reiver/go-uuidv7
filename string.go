package uuidv7

// String return the string version of a [16]byte UUID.
//
// Example usage:
//
//	uuid := uuidv7.String(uuidv7.Generate())
func String(uuid [16]byte) string {
	var hexchars string = "0123456789abcdef"

	// 16 bytes -> 32 characters + 4 hyphens = 36 characters total
	var buffer [36]byte

	buffer[ 0] = hexchars[uuid[0]>>4]
	buffer[ 1] = hexchars[uuid[0]&0x0f]
	buffer[ 2] = hexchars[uuid[1]>>4]
	buffer[ 3] = hexchars[uuid[1]&0x0f]
	buffer[ 4] = hexchars[uuid[2]>>4]
	buffer[ 5] = hexchars[uuid[2]&0x0f]
	buffer[ 6] = hexchars[uuid[3]>>4]
	buffer[ 7] = hexchars[uuid[3]&0x0f]
	buffer[ 8] = '-'
	buffer[ 9] = hexchars[uuid[4]>>4]
	buffer[10] = hexchars[uuid[4]&0x0f]
	buffer[11] = hexchars[uuid[5]>>4]
	buffer[12] = hexchars[uuid[5]&0x0f]
	buffer[13] = '-'
	buffer[14] = hexchars[uuid[6]>>4]
	buffer[15] = hexchars[uuid[6]&0x0f]
	buffer[16] = hexchars[uuid[7]>>4]
	buffer[17] = hexchars[uuid[7]&0x0f]
	buffer[18] = '-'
	buffer[19] = hexchars[uuid[8]>>4]
	buffer[20] = hexchars[uuid[8]&0x0f]
	buffer[21] = hexchars[uuid[9]>>4]
	buffer[22] = hexchars[uuid[9]&0x0f]
	buffer[23] = '-'
	buffer[24] = hexchars[uuid[10]>>4]
	buffer[25] = hexchars[uuid[10]&0x0f]
	buffer[26] = hexchars[uuid[11]>>4]
	buffer[27] = hexchars[uuid[11]&0x0f]
	buffer[28] = hexchars[uuid[12]>>4]
	buffer[29] = hexchars[uuid[12]&0x0f]
	buffer[30] = hexchars[uuid[13]>>4]
	buffer[31] = hexchars[uuid[13]&0x0f]
	buffer[32] = hexchars[uuid[14]>>4]
	buffer[33] = hexchars[uuid[14]&0x0f]
	buffer[34] = hexchars[uuid[15]>>4]
	buffer[35] = hexchars[uuid[15]&0x0f]

	return string(buffer[:])
}
