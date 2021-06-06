package qrcode

type Mode int

const (
	Numeric      Mode = 1
	Alphanumeric Mode = 2
	Byte         Mode = 4
	Kanji        Mode = 8
)
