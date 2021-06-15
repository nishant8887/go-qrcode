package qrcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	tests := []struct {
		row             int
		column          int
		expectedResults []int
	}{
		{
			row:             0,
			column:          0,
			expectedResults: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			row:             1,
			column:          0,
			expectedResults: []int{1, 1, 0, 1, 0, 0, 0, 1, 0},
		},
		{
			row:             0,
			column:          1,
			expectedResults: []int{1, 0, 1, 1, 0, 0, 0, 1, 0},
		},
		{
			row:             2,
			column:          0,
			expectedResults: []int{0, 0, 0, 2, 1, 0, 0, 0, 0},
		},
		{
			row:             2,
			column:          1,
			expectedResults: []int{1, 0, 1, 0, 1, 2, 0, 1, 0},
		},
		{
			row:             0,
			column:          2,
			expectedResults: []int{0, 0, 2, 2, 0, 0, 0, 0, 0},
		},
		{
			row:             1,
			column:          2,
			expectedResults: []int{1, 1, 2, 0, 0, 2, 0, 1, 0},
		},
		{
			row:             3,
			column:          0,
			expectedResults: []int{1, 1, 0, 0, 1, 0, 0, 1, 0},
		},
		{
			row:             3,
			column:          1,
			expectedResults: []int{0, 1, 1, 1, 1, 1, 1, 0, 0},
		},
		{
			row:             3,
			column:          2,
			expectedResults: []int{1, 1, 2, 2, 1, 0, 0, 1, 0},
		},
		{
			row:             0,
			column:          3,
			expectedResults: []int{1, 0, 0, 0, 1, 0, 0, 1, 0},
		},
		{
			row:             1,
			column:          3,
			expectedResults: []int{0, 1, 0, 1, 1, 1, 1, 0, 0},
		},
		{
			row:             2,
			column:          3,
			expectedResults: []int{1, 0, 0, 2, 0, 0, 0, 1, 0},
		},
		{
			row:             3,
			column:          4,
			expectedResults: []int{1, 1, 1, 1, 0, 0, 0, 1, 0},
		},
		{
			row:             4,
			column:          3,
			expectedResults: []int{1, 0, 0, 1, 1, 0, 0, 1, 0},
		},
	}

	for j, test := range tests {
		t.Run(fmt.Sprintf("test #%d", j), func(t *testing.T) {
			r := []int{}
			for i := 0; i < 9; i++ {
				r = append(r, mask(test.row, test.column, i))
			}
			assert.Equal(t, test.expectedResults, r, "expected to be equal")
		})
	}
}
