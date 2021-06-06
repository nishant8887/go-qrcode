package qrcode

var (
	numeric      = []rune("0123456789")
	alphanumeric = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")
)

func IsDigit(r rune) bool {
	for _, l := range numeric {
		if r == l {
			return true
		}
	}
	return false
}

func IsLetter(r rune) bool {
	for _, l := range alphanumeric {
		if r == l {
			return true
		}
	}
	return false
}

func AlphanumericIndex(r rune) int {
	for index, l := range numeric {
		if r == l {
			return index
		}
	}
	for index, l := range alphanumeric {
		if r == l {
			return 10 + index
		}
	}
	return -1
}
