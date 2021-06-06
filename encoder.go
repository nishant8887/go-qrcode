package qrcode

type rsEncoder struct {
	antilog   map[int]int
	log       map[int]int
	generator []byte
}

func NewRSEncoder(errorCodewords int) *rsEncoder {
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
	return &rsEncoder{antilog: antilog, log: log, generator: f}
}

func (e *rsEncoder) Encode(block []byte) []byte {
	p := make([]byte, len(e.generator)-1)
	p = append(block, p...)
	return e.divide(p)
}

func (e *rsEncoder) divide(cp []byte) []byte {
	if len(cp) < len(e.generator) {
		return cp
	}

	for cp[0] == 0 {
		cp = cp[1:]
	}

	mp := make([]byte, len(cp)-len(e.generator))
	mp = append(e.generator, mp...)

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
