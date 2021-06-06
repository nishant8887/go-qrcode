package qrcode

// RSEncoder - Reed-solomon encoder for QR code
type RSEncoder struct {
	antilog   map[int]int
	log       map[int]int
	generator []byte
}

// NewRSEncoder - New reed-solomon encoder
func NewRSEncoder(errorCodewords int) *RSEncoder {
	antilog := make(map[int]int)
	log := make(map[int]int)

	antilog[0] = 1
	for i := 1; i < 256; i++ {
		v := antilog[i-1] * 2
		if v >= 256 {
			v = v ^ 285
		}
		antilog[i] = v
	}

	for k, v := range antilog {
		if v == 1 && k == 255 {
			continue
		}
		log[v] = k
	}

	// Generator polynomial
	f := []byte{1}
	for i := 0; i < errorCodewords; i++ {
		size := len(f) + 1
		r := make([]byte, size)
		for j := size - 1; j > 0; j-- {
			cf := f[j-1]
			if cf != 0 {
				r[j] ^= byte(antilog[(i+log[int(cf)])%255])
			}
			r[j-1] ^= cf
		}
		f = r
	}
	return &RSEncoder{antilog: antilog, log: log, generator: f}
}

// Encode - encodes the code word to generate reed-solomon error code word
func (e *RSEncoder) Encode(block []byte) ([]byte, error) {
	p := make([]byte, len(block)+len(e.generator)-1)
	copy(p, block)
	return e.divide(p), nil
}

func (e *RSEncoder) divide(cp []byte) []byte {
	for len(cp) > 0 && cp[0] == 0 {
		cp = cp[1:]
	}

	if len(cp) < len(e.generator) {
		return cp
	}

	mp := make([]byte, len(cp))
	copy(mp, e.generator)

	for i := 0; i < len(mp); i++ {
		if mp[i] == 0 {
			continue
		}
		mp[i] = byte(e.antilog[(e.log[int(mp[i])]+e.log[int(cp[0])])%255])
	}

	r := make([]byte, len(mp)-1)
	for i := 0; i < len(r); i++ {
		r[i] = mp[i+1] ^ cp[i+1]
	}

	return e.divide(r)
}
