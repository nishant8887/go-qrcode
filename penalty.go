package qrcode

import (
	"math"
)

func penaltyForMatrix(matrix [][]bool) int {
	penalty := 0
	penalty += penaltyOneH(matrix)
	penalty += penaltyOneV(matrix)
	penalty += penaltyTwo(matrix)
	penalty += penaltyThree(matrix)
	penalty += penaltyFour(matrix)
	return penalty
}

func penaltyOneH(matrix [][]bool) int {
	penalty := 0
	size := len(matrix)
	for j := 0; j < size; j++ {
		istart := 0
		color := -1

		i := 0
		for i < size {
			pixelColor := 0
			if matrix[j][i] {
				pixelColor = 1
			}
			if color != pixelColor {
				if (i - istart) >= 5 {
					penalty += 3 + (i - istart) - 5
				}
				istart = i
				color = pixelColor
			}
			i++
		}

		if (i - istart) >= 5 {
			penalty += 3 + (i - istart) - 5
		}
	}
	return penalty
}

func penaltyOneV(matrix [][]bool) int {
	penalty := 0
	size := len(matrix)
	for j := 0; j < size; j++ {
		istart := 0
		color := -1

		i := 0
		for i < size {
			pixelColor := 0
			if matrix[i][j] {
				pixelColor = 1
			}
			if color != pixelColor {
				if (i - istart) >= 5 {
					penalty += 3 + (i - istart) - 5
				}
				istart = i
				color = pixelColor
			}
			i++
		}

		if (i - istart) >= 5 {
			penalty += 3 + (i - istart) - 5
		}
	}
	return penalty
}

func penaltyTwo(matrix [][]bool) int {
	penalty := 0
	size := len(matrix)
	for j := 0; j < size-1; j++ {
		for i := 0; i < size-1; i++ {
			if matrix[j][i] && matrix[j+1][i] && matrix[j][i+1] && matrix[i+1][j+1] ||
				!matrix[j][i] && !matrix[j+1][i] && !matrix[j][i+1] && !matrix[i+1][j+1] {
				penalty += 3
			}
		}
	}
	return penalty
}

func penaltyThree(matrix [][]bool) int {
	penalty := 0
	size := len(matrix)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if i < size-10 {
				if matrix[j][i] && !matrix[j][i+1] && matrix[j][i+2] && matrix[j][i+3] &&
					matrix[j][i+4] && !matrix[j][i+5] && matrix[j][i+6] && !matrix[j][i+7] &&
					!matrix[j][i+8] && !matrix[j][i+9] && !matrix[j][i+10] {
					penalty += 40
				}

				if !matrix[j][i] && !matrix[j][i+1] && !matrix[j][i+2] && !matrix[j][i+3] &&
					matrix[j][i+4] && !matrix[j][i+5] && matrix[j][i+6] && matrix[j][i+7] &&
					matrix[j][i+8] && !matrix[j][i+9] && matrix[j][i+10] {
					penalty += 40
				}
			}

			if j < size-10 {
				if matrix[j][i] && !matrix[j+1][i] && matrix[j+2][i] && matrix[j+3][i] &&
					matrix[j+4][i] && !matrix[j+5][i] && matrix[j+6][i] && !matrix[j+7][i] &&
					!matrix[j+8][i] && !matrix[j+9][i] && !matrix[j+10][i] {
					penalty += 40
				}

				if !matrix[j][i] && !matrix[j+1][i] && !matrix[j+2][i] && !matrix[j+3][i] &&
					matrix[j+4][i] && !matrix[j+5][i] && matrix[j+6][i] && matrix[j+7][i] &&
					matrix[j+8][i] && !matrix[j+9][i] && matrix[j+10][i] {
					penalty += 40
				}
			}
		}
	}
	return penalty
}

func penaltyFour(matrix [][]bool) int {
	size := len(matrix)
	ones := 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if matrix[j][i] {
				ones++
			}
		}
	}

	percentOnes := int(ones * 100 / (size * size))

	var l, u int
	if percentOnes%5 == 0 {
		l = (percentOnes - 5) - 50
		u = (percentOnes + 5) - 50
	} else {
		l = int(math.Floor(float64(percentOnes)/5))*5 - 50
		u = int(math.Ceil(float64(percentOnes)/5))*5 - 50
	}

	if l < 0 {
		l = -l
	}

	if u < 0 {
		u = -u
	}

	if l < u {
		return l * 10
	}
	return u * 10
}
