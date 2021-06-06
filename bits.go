package qrcode

import (
	"bytes"
	"errors"
)

var (
	errSizeExceeded = errors.New("size exceeded")
)

type bitsBuffer struct {
	buffer []bool
}

func (b *bitsBuffer) WriteInt(n int, size int) error {
	bits := make([]bool, size)
	i := 1
	for n >= 1 {
		if size-i < 0 {
			return errSizeExceeded
		}
		bits[size-i] = (n%2 == 1)
		i++
		n = n / 2
	}
	b.buffer = append(b.buffer, bits...)
	return nil
}

func (b *bitsBuffer) WriteInt64(n int64, size int) error {
	bits := make([]bool, size)
	i := 1
	for n >= 1 {
		if size-i < 0 {
			return errSizeExceeded
		}
		bits[size-i] = (n%2 == 1)
		i++
		n = n / 2
	}
	b.buffer = append(b.buffer, bits...)
	return nil
}

func (b *bitsBuffer) WriteByte(n byte) error {
	for i := 7; i >= 0; i-- {
		var bit byte
		if i > 0 {
			bit = n & byte(2<<(i-1))
		} else {
			bit = n & byte(1)
		}
		b.buffer = append(b.buffer, bit != 0)
	}
	return nil
}

func (b *bitsBuffer) Bytes() ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	size := len(b.buffer)
	for i := 0; i < size; i += 8 {
		var (
			c   *byte
			err error
		)
		if i > size-8 {
			c, err = getByte(b.buffer[i:])
		} else {
			c, err = getByte(b.buffer[i : i+8])
		}
		if err != nil {
			return nil, err
		}
		byteBuffer.WriteByte(*c)
	}
	return byteBuffer.Bytes(), nil
}

func (b *bitsBuffer) Bits() []bool {
	return b.buffer
}

func (b *bitsBuffer) Size() int {
	return len(b.buffer)
}

func (b *bitsBuffer) String() string {
	bitStr := ""
	for i := 0; i < len(b.buffer); i++ {
		if b.buffer[i] {
			bitStr += "1"
		} else {
			bitStr += "0"
		}
	}
	return bitStr
}

func getByte(n []bool) (*byte, error) {
	if len(n) > 8 {
		return nil, errors.New("length should be less than 8 for a byte")
	}
	var r byte
	for i, v := range n {
		if v {
			if i != 7 {
				r |= byte(2 << (7 - i - 1))
			} else {
				r |= byte(1)
			}
		}
	}
	return &r, nil
}
