package algorithms_test

import (
	"testing"

	. "github.com/poltao/milk/algorithms"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 6, 2, 1}, []int{1, 2, 3, 6}},
		{[]int{5}, []int{5}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := QuickSort(tc.input)
		AssertEqual(t, result, tc.expected)
	}
}
