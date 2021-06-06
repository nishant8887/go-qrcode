package qrcode

import (
	"math"
)

var finderPattern = [][]byte{
	{9, 9, 9, 9, 9, 9, 9},
	{9, 8, 8, 8, 8, 8, 9},
	{9, 8, 9, 9, 9, 8, 9},
	{9, 8, 9, 9, 9, 8, 9},
	{9, 8, 9, 9, 9, 8, 9},
	{9, 8, 8, 8, 8, 8, 9},
	{9, 9, 9, 9, 9, 9, 9},
}

var alignmentPattern = [][]byte{
	{9, 9, 9, 9, 9},
	{9, 8, 8, 8, 9},
	{9, 8, 9, 8, 9},
	{9, 8, 8, 8, 9},
	{9, 9, 9, 9, 9},
}

func (c *QRCode) generate() error {
	size := c.Size()
	matrix := make([][]byte, size)
	for i := range matrix {
		matrix[i] = make([]byte, size)
	}

	// Add finder patterns
	putMatrix(matrix, finderPattern, 0, 0)
	putMatrix(matrix, finderPattern, 0, size-len(finderPattern))
	putMatrix(matrix, finderPattern, size-len(finderPattern), 0)

	// Add finding pattern buffer
	putMatrixValue(matrix, 7, 0, 1, 8, 8)
	putMatrixValue(matrix, size-8, 0, 1, 8, 8)
	putMatrixValue(matrix, 7, size-8, 1, 8, 8)

	// Add finding pattern buffer
	putMatrixValue(matrix, 0, 7, 8, 1, 8)
	putMatrixValue(matrix, 0, size-8, 8, 1, 8)
	putMatrixValue(matrix, size-8, 7, 8, 1, 8)

	// Add alignment patters
	if c.version > 1 {
		alignmentCenters := c.alignmentPatterns()
		for _, i := range alignmentCenters {
			for _, j := range alignmentCenters {
				if (i <= 7 && (j <= 7 || j >= size-8)) || (i >= size-8 && j <= 7) {
					continue
				}
				putMatrix(matrix, alignmentPattern, i-2, j-2)
			}
		}
	}

	// Add timing patterns
	for i := 8; i < size-8; i++ {
		matrix[i][6] = byte(8 + (i+1)%2)
		matrix[6][i] = byte(8 + (i+1)%2)
	}

	// Add space for format information
	putMatrixValue(matrix, 8, 0, 1, 9, 7)
	putMatrixValue(matrix, 8, size-8, 1, 8, 7)
	putMatrixValue(matrix, 0, 8, 9, 1, 7)
	putMatrixValue(matrix, size-8, 8, 8, 1, 7)

	// Add space for version information
	if c.version >= 7 {
		putMatrixValue(matrix, size-11, 0, 3, 6, 7)
		putMatrixValue(matrix, 0, size-11, 6, 3, 7)
	}

	index := 0
	up := true
	x := size - 1
	for x > 0 {
		for y := 0; y < size; y++ {
			for i := 0; i < 2; i++ {
				if x-i == 6 {
					x = x - 1
				}
				mIndex := y
				if up {
					mIndex = size - 1 - y
				}
				if matrix[mIndex][x-i] == 0 {
					mBit := byte(0)
					if c.bitData[index] {
						mBit = byte(1)
					}
					matrix[mIndex][x-i] = mBit
					index++
				}
			}
		}
		up = !up
		x = x - 2
	}

	var (
		minPenaltyMatrix [][]bool
		minPenaltyMask   int
	)

	minPenalty := math.MaxInt32
	for mask := 0; mask < 8; mask++ {
		mv := maskedVersion(matrix, mask)
		mp := penaltyForMatrix(mv)
		if mp < minPenalty {
			minPenaltyMatrix = mv
			minPenaltyMask = mask
			minPenalty = mp
		}
	}

	c.mask = minPenaltyMask

	formatBits, err := c.formatInformation()
	if err != nil {
		return err
	}

	i := 0
	j := 0
	for k := 0; k < 7; k++ {
		if minPenaltyMatrix[8][i] {
			i++
		}
		minPenaltyMatrix[8][i] = formatBits[k]
		minPenaltyMatrix[size-1-j][8] = formatBits[k]
		i++
		j++
	}

	i = 0
	j = 0
	for k := 7; k < 15; k++ {
		if minPenaltyMatrix[8-i][8] {
			i++
		}
		minPenaltyMatrix[8-i][8] = formatBits[k]
		minPenaltyMatrix[8][size-8+j] = formatBits[k]
		i++
		j++
	}

	if c.version >= 7 {
		versionBits, err := c.versionInformation()
		if err != nil {
			return err
		}

		for j := 0; j < 6; j++ {
			for i := 0; i < 3; i++ {
				k := j*3 + i
				minPenaltyMatrix[size-11+i][j] = versionBits[17-k]
				minPenaltyMatrix[j][size-11+i] = versionBits[17-k]
			}
		}
	}

	c.code = minPenaltyMatrix
	return nil
}

func (c *QRCode) alignmentPatterns() []int {
	s := 21 + (c.version-1)*4
	d := int(c.version/7) + 1
	interval := int(math.Ceil(float64(s-13) / float64(d)))
	if interval%2 != 0 {
		interval += 1
	}
	v := make([]int, d+1)
	v[0] = 6
	v[d] = s - 7
	for i := d - 1; i > 0; i-- {
		v[i] = v[i+1] - interval
	}
	return v
}

func maskedVersion(matrix [][]byte, maskType int) [][]bool {
	size := len(matrix)
	masked := make([][]bool, size)
	for i := range masked {
		masked[i] = make([]bool, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[j][i] == 1 || matrix[j][i] == 0 {
				if mask(j, i, maskType) == 0 {
					masked[j][i] = xor(true, matrix[j][i] == 1)
				} else {
					masked[j][i] = matrix[j][i] == 1
				}
			} else if matrix[j][i] == 9 {
				masked[j][i] = true
			} else if matrix[j][i] == 8 {
				masked[j][i] = false
			}
		}
	}
	masked[size-8][8] = true
	return masked
}

func putMatrix(matrix [][]byte, m [][]byte, x, y int) {
	for j, row := range m {
		for i, column := range row {
			if matrix[y+i][x+j] == 0 {
				matrix[y+i][x+j] = column
			}
		}
	}
}

func putMatrixValue(matrix [][]byte, x, y, w, h int, value byte) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[y+i][x+j] == 0 {
				matrix[y+i][x+j] = value
			}
		}
	}
}

func mask(row, column, maskType int) int {
	switch maskType {
	case 0:
		return (row + column) % 2
	case 1:
		return row % 2
	case 2:
		return column % 3
	case 3:
		return (row + column) % 3
	case 4:
		return (int(row/2) + int(column/3)) % 2
	case 5:
		return (row*column)%2 + (row*column)%3
	case 6:
		return ((row*column)%2 + (row*column)%3) % 2
	case 7:
		return ((row+column)%2 + (row*column)%3) % 2
	}
	return 0
}
