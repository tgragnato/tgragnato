package pilsung

// Multiply by x modulo x^8 + x^4 + x^3 + x + 1
func xtime(b byte) byte {
	if b&0x80 != 0 {
		return (b << 1) ^ 0x1b
	}
	return b << 1
}

// GF(2^8) generic multiplication, double-and-add
func multiply(a byte, b byte) byte {
	var c byte = 0x0
	for i := 0; i < 8; i++ {
		if ((b >> i) & 1) != 0 {
			c ^= a
		}
		a = xtime(a)
	}
	return c
}

// Rotate Left
func rotateLeft(x byte, c int) byte {
	return (x << c % 8) | (x >> (8 - c) % 8)
}

// Unmodified AES S-box
func subByte(x0 byte) byte {
	x1 := multiply(x0, x0)    // x^2
	x2 := multiply(x1, x0)    // x^3
	x3 := multiply(x2, x2)    // x^6
	x4 := multiply(x3, x3)    // x^12
	x5 := multiply(x4, x2)    // x^15
	x6 := multiply(x5, x5)    // x^30
	x7 := multiply(x6, x6)    // x^60
	x8 := multiply(x7, x2)    // x^63
	x9 := multiply(x8, x8)    // x^126
	x10 := multiply(x9, x0)   // x^127
	x11 := multiply(x10, x10) // x^254 = x^-1
	return x11 ^ rotateLeft(x11, 1) ^ rotateLeft(x11, 2) ^
		rotateLeft(x11, 3) ^ rotateLeft(x11, 4) ^ 0x63
}

// Rotate the word
func rotWord(w [4]byte) [4]byte {
	t := w[0]
	w[0] = w[1]
	w[1] = w[2]
	w[2] = w[3]
	w[3] = t
	return w
}

// Subtract the word
func subWord(w [4]byte) [4]byte {
	for i := 0; i < 4; i++ {
		w[i] = subByte(w[i])
	}
	return w
}

// Sort array p of size n according to 0-1 array s
// Assumes s has n / 2 zeros and n / 2 ones
func getOne(s []byte, p []byte) []byte {
	a := 0
	b := 0
	buf := p
	for i := 0; i < len(p); i++ {
		if s[i] != 0x00 {
			buf[len(p)/2+a] = p[i]
			a++
		} else {
			buf[b] = p[i]
			b++
		}
	}
	return p
}

type blockT [4][4]byte

// Mix block columns
func (block *blockT) mixColumns() {
	for i := 0; i < 4; i++ {
		c0 := block[i][0]
		c1 := block[i][1]
		c2 := block[i][2]
		c3 := block[i][3]
		block[i][0] = xtime(c0^c1) ^ c1 ^ c2 ^ c3
		block[i][1] = c0 ^ xtime(c1^c2) ^ c2 ^ c3
		block[i][2] = c0 ^ c1 ^ xtime(c2^c3) ^ c3
		block[i][3] = xtime(c0^c3) ^ c0 ^ c1 ^ c2
	}
}

// Add round key to block
func (block *blockT) addRoundKey(k [4][4]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			block[i][j] ^= k[i][j]
		}
	}
}

// Sub bytes to block
func (block *blockT) subBytes(ctx *pilsungCtx, round int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			block[i][j] = ctx.sboxes[round][i][j][block[i][j]]
		}
	}
}

// Shift block rows
func (block *blockT) shiftRows(ctx *pilsungCtx, round int) {
	var copy [16]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			copy[i*4+j] = block[i][j]
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			block[i][j] = copy[ctx.pboxes[round][i*4+j]]
		}
	}
}
