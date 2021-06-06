package qrcode

var (
	formatGenerator  = []bool{true, false, true, false, false, true, true, false, true, true, true}
	versionGenerator = []bool{true, true, true, true, true, false, false, true, false, false, true, false, true}
	formatMask       = []bool{true, false, true, false, true, false, false, false, false, false, true, false, false, true, false}
)

func (c *QRCode) formatInformation() ([]bool, error) {
	buf := bitsBuffer{}

	err := buf.WriteInt(int(c.ecl), 2)
	if err != nil {
		return nil, err
	}

	err = buf.WriteInt(c.mask, 3)
	if err != nil {
		return nil, err
	}

	formatBits := golayEncode(buf.Bits(), formatGenerator)
	for i := 0; i < len(formatBits); i++ {
		formatBits[i] = xor(formatBits[i], formatMask[i])
	}

	return formatBits, nil
}

func (c *QRCode) versionInformation() ([]bool, error) {
	buf := bitsBuffer{}

	err := buf.WriteInt(c.version, 6)
	if err != nil {
		return nil, err
	}

	versionBits := golayEncode(buf.Bits(), versionGenerator)
	return versionBits, nil
}

func golayEncode(bits []bool, gp []bool) []bool {
	cp := append(bits, make([]bool, len(gp)-1)...)
	for {
		for !cp[0] {
			cp = cp[1:]
		}

		if len(cp) < len(gp) {
			break
		}

		r := make([]bool, len(cp)-1)
		m := append(gp, make([]bool, len(cp)-len(gp))...)
		for i := 0; i < len(r); i++ {
			r[i] = xor(m[i+1], cp[i+1])
		}
		cp = r
	}

	if len(cp) < len(gp) {
		cp = append(make([]bool, (len(gp)-len(cp)-1)), cp...)
	}

	return append(bits, cp...)
}

func xor(a, b bool) bool {
	return (!a && b) || (a && !b)
}
