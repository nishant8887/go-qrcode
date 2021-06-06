package qrcode

// Mode is the mode according to the QR code specifications
type Mode int

// Numeric Alphanumeric Byte Kanji modes
const (
	Numeric      Mode = 1
	Alphanumeric Mode = 2
	Byte         Mode = 4
	Kanji        Mode = 8
)
