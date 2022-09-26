package pilsung

import "crypto/sha1"

// Pilsung tweaked hashing function
func shaSign(in []byte) []byte {

	var out = []byte{}
	for i := 0; i < len(in); i += 16 {

		// Pass 16 bytes through SHA-1
		blocklen := 16
		if i+blocklen > len(in) {
			blocklen = len(in) - i
		}
		digest := sha1.Sum(in[i : i+blocklen])

		// Pilsung's tweak to SHA-1
		for i := 0; i < 20; i += 4 {
			digest[i+3] ^= 0xFF
		}
		for i := 0; i < blocklen; i++ {
			out = append(out, digest[i])
		}
	}

	return out
}

// This is the KDF
func shaKey(in []byte, outlen int) []byte {
	var out = []byte{}
	if len(in) <= outlen {
		return shaSign(in)
	}

	// Limit the dimension of the key
	maxlength := 256
	if len(in) < 256 {
		maxlength = len(in)
	}

	// Pass the chunk though the custom hash function
	for i := 0; i < maxlength; i += 32 {
		chunklen := 32
		if i+chunklen > maxlength {
			chunklen = maxlength - i
		}
		buffer := shaSign(in[i : i+chunklen])
		for j := 0; j < chunklen; j++ {
			out[i+j] = in[i+j] ^ buffer[j]
		}
	}

	return out
}

// This is the AES key schedule
func (ctx *pilsungCtx) expandRoundkey(k []byte, Nr int) {
	var (
		Nk   int  = Nr - 6 // 5
		rcon byte = 1
	)

	for i := 0; i < Nk; i++ {
		for j := 0; j < 4; j++ {
			ctx.eKey[i*Nk+j] = k[4*i+j]
		}
	}

	for i := Nk; i < 4*(Nr+1); i++ {
		var temp [4]byte
		for j := 0; j < 4; j++ {
			temp[j] = ctx.eKey[(i-1)*Nk+j]
		}
		if i%Nk == 0 {
			temp = rotWord(temp)
			temp = subWord(temp)
			temp[0] ^= rcon
			rcon = xtime(rcon)
		} else if Nk > 6 && i%Nk == 4 {
			temp = subWord(temp)
		}
		for j := 0; j < 4; j++ {
			ctx.eKey[i*Nk+j] = ctx.eKey[(i-Nk)*Nk+j] ^ temp[j]
		}
	}
}

// Key schedule wrapper
func (ctx *pilsungCtx) expandKey(k []byte) {
	derived := shaKey(k, gCryptoKeyLen)
	ctx.expandRoundkey(derived, gRoundCnt)
}

// Get the key for the current round in the right format
func (ctx *pilsungCtx) formatKey(round int) [4][4]byte {
	var output [4][4]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			output[i][j] = ctx.eKey[16*round+4*i+j]
		}
	}
	return output
}
