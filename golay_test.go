package qrcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatInformation(t *testing.T) {
	solutions := [][]bool{
		{true, true, true, false, true, true, true, true, true, false, false, false, true, false, false},
		{true, true, true, false, false, true, false, true, true, true, true, false, false, true, true},
		{true, true, true, true, true, false, true, true, false, true, false, true, false, true, false},
		{true, true, true, true, false, false, false, true, false, false, true, true, true, false, true},
		{true, true, false, false, true, true, false, false, false, true, false, true, true, true, true},
		{true, true, false, false, false, true, true, false, false, false, true, true, false, false, false},
		{true, true, false, true, true, false, false, false, true, false, false, false, false, false, true},
		{true, true, false, true, false, false, true, false, true, true, true, false, true, true, false},
		{true, false, true, false, true, false, false, false, false, false, true, false, false, true, false},
		{true, false, true, false, false, false, true, false, false, true, false, false, true, false, true},
		{true, false, true, true, true, true, false, false, true, true, true, true, true, false, false},
		{true, false, true, true, false, true, true, false, true, false, false, true, false, true, true},
		{true, false, false, false, true, false, true, true, true, true, true, true, false, false, true},
		{true, false, false, false, false, false, false, true, true, false, false, true, true, true, false},
		{true, false, false, true, true, true, true, true, false, false, true, false, true, true, true},
		{true, false, false, true, false, true, false, true, false, true, false, false, false, false, false},
		{false, true, true, false, true, false, true, false, true, false, true, true, true, true, true},
		{false, true, true, false, false, false, false, false, true, true, false, true, false, false, false},
		{false, true, true, true, true, true, true, false, false, true, true, false, false, false, true},
		{false, true, true, true, false, true, false, false, false, false, false, false, true, true, false},
		{false, true, false, false, true, false, false, true, false, true, true, false, true, false, false},
		{false, true, false, false, false, false, true, true, false, false, false, false, false, true, true},
		{false, true, false, true, true, true, false, true, true, false, true, true, false, true, false},
		{false, true, false, true, false, true, true, true, true, true, false, true, true, false, true},
		{false, false, true, false, true, true, false, true, false, false, false, true, false, false, true},
		{false, false, true, false, false, true, true, true, false, true, true, true, true, true, false},
		{false, false, true, true, true, false, false, true, true, true, false, false, true, true, true},
		{false, false, true, true, false, false, true, true, true, false, true, false, false, false, false},
		{false, false, false, false, true, true, true, false, true, true, false, false, false, true, false},
		{false, false, false, false, false, true, false, false, true, false, true, false, true, false, true},
		{false, false, false, true, true, false, true, false, false, false, false, true, true, false, false},
		{false, false, false, true, false, false, false, false, false, true, true, true, false, true, true},
	}
	ecls := []Ecl{L, M, Q, H}
	for i, ecl := range ecls {
		for mask := 0; mask < 8; mask++ {
			t.Run(fmt.Sprintf("ecl %d and mask %d", ecl, mask), func(t *testing.T) {
				qrcode := QRCode{ecl: ecl, mask: mask}
				formatInfo, err := qrcode.formatInformation()
				assert.Nil(t, err)
				assert.Equal(t, solutions[i*8+mask], formatInfo, "expected to be equal")
			})
		}
	}
}

func TestVersionInformation(t *testing.T) {
	solutions := [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true},
		{false, false, false, false, true, false, false, false, false, true, false, true, true, false, true, true, true, true},
		{false, false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true, false},
		{false, false, false, true, false, false, false, false, true, false, true, true, false, true, true, true, true, false},
		{false, false, false, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true},
		{false, false, false, true, true, false, false, false, true, true, true, false, true, true, false, false, false, true},
		{false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true, false, false},
		{false, false, true, false, false, false, false, true, false, true, true, false, true, true, true, true, false, false},
		{false, false, true, false, false, true, true, false, true, false, true, false, false, true, true, false, false, true},
		{false, false, true, false, true, false, false, true, false, false, true, true, false, true, false, false, true, true},
		{false, false, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, false},
		{false, false, true, true, false, false, false, true, true, true, false, true, true, false, false, false, true, false},
		{false, false, true, true, false, true, true, false, false, false, false, true, false, false, false, true, true, true},
		{false, false, true, true, true, false, false, true, true, false, false, false, false, false, true, true, false, true},
		{false, false, true, true, true, true, true, false, false, true, false, false, true, false, true, false, false, false},
		{false, true, false, false, false, false, true, false, true, true, false, true, true, true, true, false, false, false},
		{false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, true},
		{false, true, false, false, true, false, true, false, true, false, false, false, false, true, false, true, true, true},
		{false, true, false, false, true, true, false, true, false, true, false, false, true, true, false, false, true, false},
		{false, true, false, true, false, false, true, false, false, true, true, false, true, false, false, true, true, false},
		{false, true, false, true, false, true, false, true, true, false, true, false, false, false, false, false, true, true},
		{false, true, false, true, true, false, true, false, false, false, true, true, false, false, true, false, false, true},
		{false, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, false, false},
		{false, true, true, false, false, false, true, true, true, false, true, true, false, false, false, true, false, false},
		{false, true, true, false, false, true, false, false, false, true, true, true, true, false, false, false, false, true},
		{false, true, true, false, true, false, true, true, true, true, true, false, true, false, true, false, true, true},
		{false, true, true, false, true, true, false, false, false, false, true, false, false, false, true, true, true, false},
		{false, true, true, true, false, false, true, true, false, false, false, false, false, true, true, false, true, false},
		{false, true, true, true, false, true, false, false, true, true, false, false, true, true, true, true, true, true},
		{false, true, true, true, true, false, true, true, false, true, false, true, true, true, false, true, false, true},
		{false, true, true, true, true, true, false, false, true, false, false, true, false, true, false, false, false, false},
		{true, false, false, false, false, false, true, false, false, true, true, true, false, true, false, true, false, true},
		{true, false, false, false, false, true, false, true, true, false, true, true, true, true, false, false, false, false},
		{true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, true, false},
		{true, false, false, false, true, true, false, true, true, true, true, false, false, true, true, true, true, true},
		{true, false, false, true, false, false, true, false, true, true, false, false, false, false, true, false, true, true},
		{true, false, false, true, false, true, false, true, false, false, false, false, true, false, true, true, true, false},
		{true, false, false, true, true, false, true, false, true, false, false, true, true, false, false, true, false, false},
		{true, false, false, true, true, true, false, true, false, true, false, true, false, false, false, false, false, true},
		{true, false, true, false, false, false, true, true, false, false, false, true, true, false, true, false, false, true},
	}
	for v := 0; v <= 40; v++ {
		t.Run(fmt.Sprintf("version %d", v), func(t *testing.T) {
			qrcode := QRCode{version: v}
			versionInfo, err := qrcode.versionInformation()
			assert.Nil(t, err)
			assert.Equal(t, solutions[v], versionInfo, "expected to be equal")
		})
	}
}

func TestGolayEncode(t *testing.T) {
	tests := []struct {
		cp             []bool
		gp             []bool
		expectedResult []bool
	}{
		{
			cp:             []bool{false, false, false, false, false},
			gp:             formatGenerator,
			expectedResult: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			cp:             []bool{false, true, false, false, false},
			gp:             formatGenerator,
			expectedResult: []bool{false, true, false, false, false, true, true, true, true, false, true, false, true, true, false},
		},
		{
			cp:             []bool{true, false, false, false, false},
			gp:             formatGenerator,
			expectedResult: []bool{true, false, false, false, false, true, false, true, false, false, true, true, false, true, true},
		},
		{
			cp:             []bool{true, true, false, false, false},
			gp:             formatGenerator,
			expectedResult: []bool{true, true, false, false, false, false, true, false, true, false, false, true, true, false, true},
		},
		{
			cp:             []bool{false, true, false, false, true},
			gp:             formatGenerator,
			expectedResult: []bool{false, true, false, false, true, true, false, true, true, true, false, false, false, false, true},
		},
		{
			cp:             []bool{false, false, false, false, false, false},
			gp:             versionGenerator,
			expectedResult: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			cp:             []bool{false, false, false, false, false, true},
			gp:             versionGenerator,
			expectedResult: []bool{false, false, false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true},
		},
		{
			cp:             []bool{false, false, false, false, true, false},
			gp:             versionGenerator,
			expectedResult: []bool{false, false, false, false, true, false, false, false, false, true, false, true, true, false, true, true, true, true},
		},
		{
			cp:             []bool{false, false, false, false, true, true},
			gp:             versionGenerator,
			expectedResult: []bool{false, false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true, false},
		},
	}

	for _, test := range tests {
		result := golayEncode(test.cp, test.gp)
		assert.Equal(t, len(test.cp)+len(test.gp)-1, len(result), "expected size do not match")
		assert.Equal(t, test.expectedResult, result, "expected result do not match")
	}
}

func TestXor(t *testing.T) {
	assert.Equal(t, true, xor(true, false), "expected to be true")
	assert.Equal(t, true, xor(false, true), "expected to be true")
	assert.Equal(t, false, xor(true, true), "expected to be false")
	assert.Equal(t, false, xor(false, false), "expected to be false")
}
