package two_pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{362880, 181440, 120960, 90720, 72576, 60480, 51840, 45360, 40320}},
	}

	for _, tc := range testCases {
		got := productExceptSelf(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}

func genLargeSlice() []int {
	var nums []int
	for i := 0; i < 10000; i++ {
		nums = append(nums, i)
	}

	return nums
}

func BenchmarkProductExceptSelf(b *testing.B) {
	slice := genLargeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = productExceptSelf(slice)
	}
}

func BenchmarkProductExceptSelfNaive(b *testing.B) {
	slice := genLargeSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = productExceptSelfNaiveButFaster(slice)
	}
}
