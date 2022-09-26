package pilsung

// Pilsung Context
type pilsungCtx struct {
	sboxXorConstant     byte                // 3
	sboxes              [30][4][4][256]byte // S-Boxes
	pboxes              [30][16]byte        // P-Boxes
	currentPermutation8 [8]byte             // used for temporary storage
	eKey                [240]byte           // AES scheduled key
}

// Permute the bits of inputMask according to the computed permutation
func (ctx *pilsungCtx) getPeSb(inputMask byte) byte {
	var x byte = 0x0
	for i := 0; i < 8; i++ {
		if ((1 << i) & inputMask) != 0x0 {
			x |= 1 << ctx.currentPermutation8[i]
		}
	}
	return x ^ ctx.sboxXorConstant
}

// Generate a random permutation of [0:7] at ctx.currentPermutation8
func (ctx *pilsungCtx) getP8forSEnc(inputMask byte) {
	var (
		coinflips [24]byte
		v0        byte = treeInteger8[inputMask%70]
		v1        byte = (treeInteger4[(inputMask&0x0F)%6] << 0) | (treeInteger4[(inputMask>>4)%6] << 4)
	)

	for i := 0; i < 8; i++ {
		coinflips[i] = (v0 >> (7 - i)) & 1
	}
	for i := 0; i < 8; i++ {
		coinflips[i+8] = (v1 >> (7 - i)) & 1
	}

	for i := 0; i < 8; i += 2 {
		if (inputMask & (3 << i)) != 0x0 {
			coinflips[16+i+0] = 1
			coinflips[16+i+1] = 0
		} else {
			coinflips[16+i+0] = 0
			coinflips[16+i+1] = 1
		}
	}

	// Initialize permutation with identity
	for i := 0; i < 8; i++ {
		ctx.currentPermutation8[i] = byte(i)
	}

	// Iterative version of the Rao-Sandelius shuffle
	for i := 0; i < 3; i++ {
		bins := 1 << i
		size := 1 << (3 - i)
		for j := 0; j < bins; j++ {
			s := coinflips[i*8 : i*8+7]
			p := ctx.currentPermutation8[j*size : j*size+size-1]
			copy(ctx.currentPermutation8[:], getOne(s, p)[:])
		}
	}
}

// Generate a random permutation of [0:15] at output
func getP16Enc(input [16]byte) [16]byte {
	var (
		coinflips [64]byte
		output    [16]byte
	)

	// coin flips for first level
	for i := 0; i < 4; i++ {
		v0 := treeInteger4[(input[i]^input[i+4])%6]
		for j := 0; j < 4; j++ {
			coinflips[4*i+j] = (v0 >> (3 - j)) & 1
		}
	}

	// coin flips for second level
	for i := 0; i < 8; i++ {
		if ((input[i] >> i) & 1) != 0 {
			coinflips[16+2*i+0] = 1
			coinflips[16+2*i+1] = 0
		} else {
			coinflips[16+2*i+0] = 0
			coinflips[16+2*i+1] = 1
		}
	}

	// coin flips for third level
	for i := 0; i < 4; i++ {
		v1 := treeInteger4[(input[i+8]^input[i+12])%6]
		for j := 0; j < 4; j++ {
			coinflips[32+4*i+j] = (v1 >> (3 - j)) & 1
		}
	}

	// coin flips for fourth level
	for i := 0; i < 8; i++ {
		if ((input[8+i] >> i) & 1) != 0 {
			coinflips[48+2*i+0] = 1
			coinflips[48+2*i+1] = 0
		} else {
			coinflips[48+2*i+0] = 0
			coinflips[48+2*i+1] = 1
		}
	}

	// Initialize permutation with identity
	for i := 0; i < 16; i++ {
		output[i] = byte(i)
	}

	// Iterative version of the Rao-Sandelius shuffle
	for i := 0; i < 4; i++ {
		bins := 1 << i       // number of subgroups
		size := 1 << (4 - i) // size of each subgroup
		for j := 0; j < bins; j++ {
			s := coinflips[i*16 : i*16+15]
			p := output[j*size : j*size+size-1]
			copy(output[:], getOne(s, p)[:])
		}
	}
	return output
}

// Generate all necessary S-boxes
func (ctx *pilsungCtx) getVSboxAll() {
	for rounds := 1; rounds < gRoundCnt; rounds++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				ctx.getP8forSEnc(ctx.eKey[j+4*i+16*rounds])
				for k := 0; k < 256; k++ {
					ctx.sboxes[rounds][i][j][k] = ctx.getPeSb(subByte(byte(k)))
				}
			}
		}
	}
}

// Generate all necessary P-boxes
func (ctx *pilsungCtx) getVPboxAll() {
	for rounds := 1; rounds < gRoundCnt; rounds++ {
		var roundScheduledKey [16]byte
		copy(roundScheduledKey[:], ctx.eKey[16*rounds:16*rounds+15])
		ctx.pboxes[rounds] = getP16Enc(roundScheduledKey)
	}
}

// Initialize the xor constant
func (ctx *pilsungCtx) initVar() {
	ctx.sboxXorConstant = xorConstant
}

// Generate the encryption permutations
func (ctx *pilsungCtx) genEncPerm() {
	ctx.initVar()
	ctx.getVSboxAll()
	ctx.getVPboxAll()
}
