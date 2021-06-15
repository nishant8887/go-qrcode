package qrcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var matrixOne = [][]bool{
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{true, true, true, true, true, true, false, false, false, false, false, false},
	{false, false, false, false, false, false, true, true, true, true, true, true},
	{false, false, false, false, false, false, true, true, true, true, true, true},
	{false, false, false, false, false, false, true, true, true, true, true, true},
	{false, false, false, false, false, false, true, true, true, true, true, true},
	{false, false, false, false, false, false, true, true, true, true, true, true},
	{false, false, false, false, false, false, true, true, true, true, true, true},
}

var matrixTwo = [][]bool{
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true},
}

var matrixThree = [][]bool{
	{true, true, false, false, true, true, false, false, true, true},
	{true, true, false, false, true, true, false, false, true, true},
	{false, false, true, true, false, false, true, true, false, false},
	{false, false, true, true, false, false, true, true, false, false},
	{true, true, false, false, true, true, false, false, true, true},
	{true, true, false, false, true, true, false, false, true, true},
	{false, false, true, true, false, false, true, true, false, false},
	{false, false, true, true, false, false, true, true, false, false},
	{true, true, false, false, true, true, false, false, true, true},
	{true, true, false, false, true, true, false, false, true, true},
}

var matrixFour = [][]bool{
	{true, false, true, true, true, false, true, false, false, false, false},
	{false, false, false, false, true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true, false, false, false, false},
	{false, false, false, false, true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true, false, false, false, false},
	{false, false, false, false, true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true, false, false, false, false},
	{false, false, false, false, true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true, false, false, false, false},
	{false, false, false, false, true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true, false, false, false, false},
}

var matrixFive = [][]bool{
	{true, false, true, false, true, false, true, false, true, false, true},
	{false, false, false, false, false, false, false, false, false, false, false},
	{true, false, true, false, true, false, true, false, true, false, true},
	{true, false, true, false, true, false, true, false, true, false, true},
	{true, true, true, true, true, true, true, true, true, true, true},
	{false, false, false, false, false, false, false, false, false, false, false},
	{true, true, true, true, true, true, true, true, true, true, true},
	{false, true, false, true, false, true, false, true, false, true, false},
	{false, true, false, true, false, true, false, true, false, true, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, true, false, true, false, true, false, true, false, true, false},
}

var matrixSix = [][]bool{
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}

func TestPenaltyForMatrix(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 542,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 50,
		},
		{
			matrix:        matrixThree,
			expectedValue: 75,
		},
		{
			matrix:        matrixFour,
			expectedValue: 515,
		},
		{
			matrix:        matrixFive,
			expectedValue: 515,
		},
		{
			matrix:        matrixSix,
			expectedValue: 561,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyForMatrix(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}

func TestPenaltyOneH(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 96,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 0,
		},
		{
			matrix:        matrixThree,
			expectedValue: 0,
		},
		{
			matrix:        matrixFour,
			expectedValue: 0,
		},
		{
			matrix:        matrixFive,
			expectedValue: 45,
		},
		{
			matrix:        matrixSix,
			expectedValue: 80,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyOneH(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}

func TestPenaltyOneV(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 96,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 0,
		},
		{
			matrix:        matrixThree,
			expectedValue: 0,
		},
		{
			matrix:        matrixFour,
			expectedValue: 45,
		},
		{
			matrix:        matrixFive,
			expectedValue: 0,
		},
		{
			matrix:        matrixSix,
			expectedValue: 60,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyOneV(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}

func TestPenaltyTwo(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 300,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 0,
		},
		{
			matrix:        matrixThree,
			expectedValue: 75,
		},
		{
			matrix:        matrixFour,
			expectedValue: 30,
		},
		{
			matrix:        matrixFive,
			expectedValue: 30,
		},
		{
			matrix:        matrixSix,
			expectedValue: 171,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyTwo(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}

func TestPenaltyThree(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 0,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 0,
		},
		{
			matrix:        matrixThree,
			expectedValue: 0,
		},
		{
			matrix:        matrixFour,
			expectedValue: 440,
		},
		{
			matrix:        matrixFive,
			expectedValue: 440,
		},
		{
			matrix:        matrixSix,
			expectedValue: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyThree(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}

func TestPenaltyFour(t *testing.T) {

	tests := []struct {
		matrix        [][]bool
		expectedValue int
	}{
		{
			matrix:        matrixOne,
			expectedValue: 50,
		},
		{
			matrix:        matrixTwo,
			expectedValue: 50,
		},
		{
			matrix:        matrixThree,
			expectedValue: 0,
		},
		{
			matrix:        matrixFour,
			expectedValue: 0,
		},
		{
			matrix:        matrixFive,
			expectedValue: 0,
		},
		{
			matrix:        matrixSix,
			expectedValue: 250,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("matrix %d", i+1), func(t *testing.T) {
			penalty := penaltyFour(test.matrix)
			assert.Equal(t, test.expectedValue, penalty, "expected to be equal")
		})
	}
}
