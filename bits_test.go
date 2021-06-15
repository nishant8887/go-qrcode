package qrcode

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteInt(t *testing.T) {
	tests := []struct {
		name           string
		value          int
		size           int
		expectedResult []bool
		expectedError  error
	}{
		{
			name:           "write int with proper size",
			value:          5,
			size:           8,
			expectedResult: []bool{false, false, false, false, false, true, false, true},
			expectedError:  nil,
		},
		{
			name:           "write int with exact size",
			value:          5,
			size:           3,
			expectedResult: []bool{true, false, true},
			expectedError:  nil,
		},
		{
			name:           "write int with improper size",
			value:          8,
			size:           3,
			expectedResult: nil,
			expectedError:  errSizeExceeded,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bitsBuffer{}
			err := buf.WriteInt(test.value, test.size)

			if test.expectedError != nil {
				assert.NotNil(t, err, "error expected to be nil")
				assert.EqualError(t, test.expectedError, err.Error(), "expected error do not match")
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.expectedResult, buf.Bits(), "expected bits do not match")
				assert.Equal(t, len(test.expectedResult), buf.Size(), "expected size do not match")
			}
		})
	}
}

func TestWriteByte(t *testing.T) {
	tests := []struct {
		name           string
		data           byte
		expectedResult []bool
	}{
		{
			name:           "write byte",
			data:           5,
			expectedResult: []bool{false, false, false, false, false, true, false, true},
		},
		{
			name:           "write lowest byte",
			data:           0,
			expectedResult: []bool{false, false, false, false, false, false, false, false},
		},
		{
			name:           "write highest byte",
			data:           255,
			expectedResult: []bool{true, true, true, true, true, true, true, true},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bitsBuffer{}
			err := buf.WriteByte(test.data)

			assert.Equal(t, test.expectedResult, buf.Bits(), "expected bits do not match")
			assert.Equal(t, 8, buf.Size(), "expected size do not match")
			assert.Nil(t, err)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		data           []bool
		expectedResult string
	}{
		{
			data:           []bool{false, false, false, false},
			expectedResult: "0000",
		},
		{
			data:           []bool{true, false, true, false, true, false, true},
			expectedResult: "1010101",
		},
		{
			data:           []bool{true, true, true, true, true},
			expectedResult: "11111",
		},
	}

	for _, test := range tests {
		buf := bitsBuffer{}
		buf.buffer = test.data
		assert.Equal(t, test.expectedResult, buf.String(), "expected string representations do not match")
	}
}

func TestGetByte(t *testing.T) {
	tests := []struct {
		name           string
		data           []bool
		expectedResult *byte
		expectedError  error
	}{
		{
			name:           "get byte",
			data:           []bool{false, false, false, false, false, true, false, true},
			expectedResult: bytePtr(5),
			expectedError:  nil,
		},
		{
			name:           "get byte for min value",
			data:           []bool{false, false, false, false, false, false, false, false},
			expectedResult: bytePtr(0),
			expectedError:  nil,
		},
		{
			name:           "get byte for max value",
			data:           []bool{true, true, true, true, true, true, true, true},
			expectedResult: bytePtr(255),
			expectedError:  nil,
		},
		{
			name:           "get byte for array of less than size 8",
			data:           []bool{true, true, false},
			expectedResult: bytePtr(192),
			expectedError:  nil,
		},
		{
			name:           "get byte for array of more that size 8",
			data:           []bool{false, true, true, true, true, true, true, true, true},
			expectedResult: nil,
			expectedError:  errors.New("length should be less than 8 for a byte"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := getByte(test.data)

			if test.expectedError != nil {
				assert.NotNil(t, err)
				assert.Nil(t, result)
				assert.EqualError(t, test.expectedError, err.Error(), "expected error do not match")
			} else {
				fmt.Println(*result)
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, *test.expectedResult, *result, "expected byte do not match")
			}
		})
	}
}

func TestGetBytes(t *testing.T) {
	tests := []struct {
		name           string
		data           []int
		size           []int
		expectedResult []byte
	}{
		{
			name:           "write data with multiple of 8 bits",
			data:           []int{5, 255, 136},
			size:           []int{8, 8, 8},
			expectedResult: []byte{5, 255, 136},
		},
		{
			name:           "write data with not a multiple of 8 bits",
			data:           []int{5, 255, 6},
			size:           []int{8, 8, 3},
			expectedResult: []byte{5, 255, 192},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bitsBuffer{}
			for i, v := range test.data {
				buf.WriteInt(v, test.size[i])
			}

			data, err := buf.Bytes()
			assert.Nil(t, err)
			assert.NotNil(t, data)
			assert.Equal(t, test.expectedResult, data, "expected result do not match")
		})
	}
}

func bytePtr(v byte) *byte {
	return &v
}
