package qrcode

// Mode - mode according to the QR code specifications
type Mode int

/*
Modes according to the QR code specifications

1. Numeric

2. Alphanumeric

3. Byte

4. Kanji
*/
const (
	Numeric      Mode = 1
	Alphanumeric Mode = 2
	Byte         Mode = 4
	Kanji        Mode = 8
)
