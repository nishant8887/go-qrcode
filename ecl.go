package qrcode

// Ecl is the error correction level according to the QR code specifications
type Ecl int

// L M Q H are the four levels of error correction
const (
	L Ecl = 1
	M Ecl = 0
	Q Ecl = 3
	H Ecl = 2
)
