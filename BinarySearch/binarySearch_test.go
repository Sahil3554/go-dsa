package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		result int
	}{
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 0,
			result: -1,
		},
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 4,
			result: 3,
		},
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 3,
			result: 2,
		},
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 1,
			result: 0,
		},
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 5,
			result: 4,
		},
		{
			arr:    []int{1, 2, 3, 4, 5},
			target: 2,
			result: 1,
		},
	}
	for index, tc := range testCases {
		result := binarySearch(tc.arr, tc.target)
		t.Logf("Test #%d: Searching for %d in %v. Expected: %d, Got: %d", index, tc.target, tc.arr, tc.result, result)
		assert.Equal(t, tc.result, result)
	}
}
