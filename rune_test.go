package qrcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	digitTests := []rune("0123456789")

	for _, test := range digitTests {
		assert.Equal(t, true, isDigit(test), fmt.Sprintf("%c expected a digit", test))
	}

	nonDigitTests := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()[]{};:'\"?/<,>.|\\-_+=` ")
	for _, test := range nonDigitTests {
		assert.Equal(t, false, isDigit(test), fmt.Sprintf("%c expected not a digit", test))
	}
}

func TestIsLetter(t *testing.T) {
	letterTests := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")

	for _, test := range letterTests {
		assert.Equal(t, true, isLetter(test), fmt.Sprintf("'%c' expected a letter", test))
	}

	nonLetterTests := []rune("abcdefghijklmnopqrstuvwxyz~!@#^&()[]{};'\"?<,>|\\_=`")
	for _, test := range nonLetterTests {
		assert.Equal(t, false, isLetter(test), fmt.Sprintf("'%c' expected not a letter", test))
	}
}

func TestAlphanumericIndex(t *testing.T) {
	tests := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")
	for index, test := range tests {
		assert.Equal(t, index, alphanumericIndex(test), fmt.Sprintf("'%c' must have %d index", test, index))
	}

	nonAlphanumericTests := []rune("abcdefghijklmnopqrstuvwxyz~!@#^&()[]{};'\"?<,>|\\_=`")
	for _, test := range nonAlphanumericTests {
		assert.Equal(t, -1, alphanumericIndex(test), fmt.Sprintf("'%c' expected index is -1", test))
	}
}
