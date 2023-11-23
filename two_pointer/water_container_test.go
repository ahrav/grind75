package two_pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	testCases := []struct {
		height []int
		want   int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 2, 1}, 2},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		got := conainterWithTheMostWater(tc.height)
		assert.Equal(t, tc.want, got)
	}
}
