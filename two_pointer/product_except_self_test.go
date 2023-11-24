package two_pointer

import "testing"

func genLargeSlice() []int {
	var nums []int
	for i := 0; i < 10000; i++ {
		nums = append(nums, i)
	}
	nums = append(nums, 0)

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
