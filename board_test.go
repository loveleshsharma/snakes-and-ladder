package main

import (
	"testing"
)

func TestNewBoard_ShouldFindCorrectPositionInBoard(t *testing.T) {
	testBoard := NewBoard()

	type testCase struct {
		pos       int
		expectedX int
		expectedY int
	}

	testCases := []testCase{
		{
			pos:       7,
			expectedX: 6,
			expectedY: 9,
		},
		{
			pos:       39,
			expectedX: 1,
			expectedY: 6,
		},
		{
			pos:       27,
			expectedX: 6,
			expectedY: 7,
		},
		{
			pos:       83,
			expectedX: 2,
			expectedY: 1,
		},
		{
			pos:       66,
			expectedX: 5,
			expectedY: 3,
		},
		{
			pos:       100,
			expectedX: 0,
			expectedY: 0,
		},
		{
			pos:       80,
			expectedX: 0,
			expectedY: 2,
		},
		{
			pos:       20,
			expectedX: 0,
			expectedY: 8,
		},
		{
			pos:       90,
			expectedX: 9,
			expectedY: 1,
		},
		{
			pos:       50,
			expectedX: 9,
			expectedY: 5,
		},
	}

	for _, eachCase := range testCases {
		x, y := testBoard.FindCoordinatesInGrid(eachCase.pos)

		if x != eachCase.expectedX || y != eachCase.expectedY {
			t.Errorf("x and y should be: %d and %d for pos: %d, actual x and y: %d and %d", eachCase.expectedX, eachCase.expectedY, eachCase.pos, x, y)
			return
		}
	}
}
