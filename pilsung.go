package pilsung

const (
	gCryptoKeyLen = 32 // 256-bit key
	gRoundCnt     = 11 // 11 rounds
	xorConstant   = 3  // S-box Xor Constant
)

var (
	// All combinations of 4 bits in 8-bit words
	treeInteger8 = [70]byte{
		0x0F, 0x87, 0x47, 0x27, 0x17, 0xC3, 0x63, 0x33, 0x1B, 0xA3,
		0x53, 0x2B, 0x93, 0x4B, 0x8B, 0xE1, 0x71, 0x39, 0x1D, 0xB1,
		0x59, 0x2D, 0x99, 0x4D, 0x8D, 0xD1, 0x69, 0x35, 0xA9, 0x55,
		0x95, 0xC9, 0x65, 0xA5, 0xC5, 0xF0, 0x78, 0x3C, 0x1E, 0xB8,
		0x5C, 0x2E, 0x9C, 0x4E, 0x8E, 0xD8, 0x6C, 0x36, 0xAC, 0x56,
		0x96, 0xCC, 0x66, 0xA6, 0xC6, 0xE8, 0x74, 0x3A, 0xB4, 0x5A,
		0x9A, 0xD4, 0x6A, 0xAA, 0xCA, 0xE4, 0x72, 0xB2, 0xD2, 0xE2,
	}
	// All combinations of 2 bits in 4-bit words
	treeInteger4 = [6]byte{
		0x03, 0x09, 0x05, 0x0C, 0x06, 0x0A,
	}
)

// Export the Encryption function
func Encrypt(input [16]byte, key []byte) [16]byte {
	var (
		block  blockT
		ctx    pilsungCtx
		output [16]byte
	)

	// Set the key
	ctx.expandKey(key)
	ctx.genEncPerm()

	// Array to matrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			block[i][j] = input[i*4+j]
		}
	}

	// First round
	block.addRoundKey(ctx.formatKey(0))

	// Middle rounds
	for round := 1; round < 10; round++ {
		block.subBytes(&ctx, round)
		block.shiftRows(&ctx, round)
		block.mixColumns()
		block.addRoundKey(ctx.formatKey(round))
	}

	// Last round
	block.subBytes(&ctx, 10)
	block.shiftRows(&ctx, 10)
	// No mixColumns
	block.addRoundKey(ctx.formatKey(10))

	// Matrix to array
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			output[i+4+j] = block[i][j]
		}
	}

	return output
}
