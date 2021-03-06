package qrcode

import (
	"errors"
	"image"
	"image/color"
	"strconv"
)

var (
	errTextTooLong      = errors.New("input too long")
	errBlocksDoNotMatch = errors.New("blocks do not match")
)

/*
QRCode - generated QRCode from the data
*/
type QRCode struct {
	version    int
	mode       Mode
	ecl        Ecl
	data       string
	headerSize int
	codeInfo   []int
	bitData    []bool
	mask       int
	code       [][]bool
}

// Size - returns the size of the QR code's grid
func (c *QRCode) Size() int {
	return 21 + (c.version-1)*4
}

// Version - returns the version of the generated QR code
func (c *QRCode) Version() int {
	return c.version
}

// Mode - returns the mode of the generated QR code
func (c *QRCode) Mode() Mode {
	return c.mode
}

// Ecl - returns the error correction level of the generated QR code
func (c *QRCode) Ecl() Ecl {
	return c.ecl
}

// Mask - returns the mask used for the generated QR code
func (c *QRCode) Mask() int {
	return c.mask
}

// Matrix - returns the 2D boolean array for the generated QR code
func (c *QRCode) Matrix() [][]bool {
	return c.code
}

// Image - returns the image of the generated QR code
func (c *QRCode) Image() image.Image {
	size := c.Size()
	codeImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	white := color.RGBA{255, 255, 255, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if c.code[y][x] {
				codeImage.Set(x, y, black)
			} else {
				codeImage.Set(x, y, white)
			}
		}
	}
	return codeImage
}

func (c *QRCode) encode() error {
	buf := bitsBuffer{}
	size := len(c.data)

	err := buf.WriteInt(int(c.mode), 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt(size, c.headerSize)
	if err != nil {
		return err
	}

	if c.mode == Numeric {
		for i := 0; i < size; i += 3 {
			var v string
			if i > size-3 {
				v = c.data[i:]
			} else {
				v = c.data[i : i+3]
			}
			n, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			if n > 99 {
				err = buf.WriteInt64(n, 10)
			} else if n > 9 {
				err = buf.WriteInt64(n, 7)
			} else {
				err = buf.WriteInt64(n, 4)
			}
			if err != nil {
				return err
			}
		}
	} else if c.mode == Alphanumeric {
		for i := 0; i < size; i += 2 {
			var v []rune
			if i > size-2 {
				v = []rune(c.data[i:])
			} else {
				v = []rune(c.data[i : i+2])
			}
			if len(v) == 2 {
				err = buf.WriteInt(alphanumericIndex(v[0])*45+alphanumericIndex(v[1]), 11)
			} else {
				err = buf.WriteInt(alphanumericIndex(v[0]), 6)
			}
			if err != nil {
				return err
			}
		}
	} else if c.mode == Byte {
		for i := 0; i < size; i++ {
			err = buf.WriteByte(c.data[i])
			if err != nil {
				return err
			}
		}
	}

	maxAllowedBits := c.codeInfo[0] * 8
	paddingBits := maxAllowedBits - buf.Size()
	if paddingBits > 4 {
		paddingBits = 4
	}

	err = buf.WriteInt(0, paddingBits)
	if err != nil {
		return err
	}

	if buf.Size()%8 != 0 {
		err = buf.WriteInt(0, 8-buf.Size()%8)
		if err != nil {
			return err
		}
	}

	totalBits := buf.Size()
	if totalBits < maxAllowedBits {
		for i := 0; i < (maxAllowedBits/8)-(totalBits/8); i++ {
			if i%2 == 0 {
				err = buf.WriteByte(236)
			} else {
				err = buf.WriteByte(17)
			}
			if err != nil {
				return err
			}
		}
	}

	codewords, err := buf.Bytes()
	if err != nil {
		return err
	}

	blocks := make([][]byte, 0)
	groupOneBlocks := c.codeInfo[2]
	groupOneBlockSize := c.codeInfo[3]
	groupTwoBlocks := c.codeInfo[4]
	groupTwoBlockSize := c.codeInfo[5]

	totalBlocks := groupOneBlocks*groupOneBlockSize + groupTwoBlocks*groupTwoBlockSize
	if len(codewords) != totalBlocks {
		return errBlocksDoNotMatch
	}

	for i := 0; i < groupOneBlocks; i++ {
		blocks = append(blocks, codewords[i*groupOneBlockSize:(i+1)*groupOneBlockSize])
	}

	for i := 0; i < groupTwoBlocks; i++ {
		blocks = append(blocks, codewords[groupOneBlocks*groupOneBlockSize+i*groupTwoBlockSize:groupOneBlocks*groupOneBlockSize+(i+1)*groupTwoBlockSize])
	}

	ecCodeBlockSize := c.codeInfo[1]
	ecCodeBlocks := make([][]byte, 0)

	rs := NewRSEncoder(ecCodeBlockSize)
	for _, block := range blocks {
		ecCodeBlock, _ := rs.Encode(block)
		ecCodeBlocks = append(ecCodeBlocks, ecCodeBlock)
	}

	buf = bitsBuffer{}
	maxBlockSize := groupOneBlockSize
	if maxBlockSize < groupTwoBlockSize {
		maxBlockSize = groupTwoBlockSize
	}
	for i := 0; i < maxBlockSize; i++ {
		for _, block := range blocks {
			if i < len(block) {
				err = buf.WriteByte(block[i])
				if err != nil {
					return err
				}
			}
		}
	}

	for i := 0; i < ecCodeBlockSize; i++ {
		for _, ecCodeBlock := range ecCodeBlocks {
			if i < len(ecCodeBlock) {
				err = buf.WriteByte(ecCodeBlock[i])
				if err != nil {
					return err
				}
			}
		}
	}

	err = buf.WriteInt(0, remainderBits[c.version-1])
	if err != nil {
		return err
	}

	c.bitData = buf.Bits()
	return nil
}

// New creates a QR code for given string and error correction level
func New(data string, ecl Ecl) (*QRCode, error) {
	mode := Numeric
	for _, r := range []rune(data) {
		if !isDigit(r) {
			mode = Alphanumeric
		}
		if !isLetter(r) {
			mode = Byte
			break
		}
	}

	eclIndex := 0
	switch ecl {
	case L:
		eclIndex = 0
	case M:
		eclIndex = 1
	case Q:
		eclIndex = 2
	case H:
		eclIndex = 3
	}

	modeIndex := 0
	switch mode {
	case Numeric:
		modeIndex = 0
	case Alphanumeric:
		modeIndex = 1
	case Byte:
		modeIndex = 2
	case Kanji:
		modeIndex = 3
	}

	version := 0
	for index, capacity := range capacities {
		if capacity[eclIndex][modeIndex] >= len(data) {
			version = index + 1
			break
		}
	}
	if version == 0 {
		return nil, errTextTooLong
	}

	headerSize := 0
	if version >= 1 && version < 10 {
		headerSize = headerSizes[0][modeIndex]
	} else if version >= 10 && version < 27 {
		headerSize = headerSizes[1][modeIndex]
	} else if version >= 27 {
		headerSize = headerSizes[2][modeIndex]
	}

	codeInfo := codewordsInfo[version-1][eclIndex]
	code := QRCode{version: version, mode: mode, ecl: ecl, data: data, headerSize: headerSize, codeInfo: codeInfo}

	err := code.encode()
	if err != nil {
		return nil, err
	}

	err = code.generate()
	if err != nil {
		return nil, err
	}

	return &code, nil
}
