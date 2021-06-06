package qrcode

// Ecl - error correction level according to the QR code specifications
type Ecl int

/*
Four error correction levels:
L - Recovers 7% of data
M - Recovers 15% of data
Q - Recovers 25% of data
H - Recovers 30% of data
*/
const (
	L Ecl = 1
	M Ecl = 0
	Q Ecl = 3
	H Ecl = 2
)
