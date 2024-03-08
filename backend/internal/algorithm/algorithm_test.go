package algorithm

import (
	"github.com/go-quicktest/qt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name      string
		target    int
		packSizes []int

		expectedPacks map[int]int
	}{
		{
			name:          "Test Case - 1",
			target:        1,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{250: 1},
		},
		{
			name:          "Test Case - 2",
			target:        250,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{250: 1},
		},
		{
			name:          "Test Case - 3",
			target:        251,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{500: 1},
		},
		{
			name:          "Test Case - 4",
			target:        501,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{500: 1, 250: 1},
		},
		{
			name:          "Test Case - 5",
			target:        12001,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, fetchedPacks := Solve(tc.target, tc.packSizes)
			qt.Assert(t, qt.IsTrue(reflect.DeepEqual(tc.expectedPacks, fetchedPacks.M)))
		})
	}
}
