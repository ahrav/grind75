package two_pointer

func productExceptSelf(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}

	res := make([]int, l)
	for i := range res {
		res[i] = 1
	}

	lProduct, rProduct := 1, 1
	lp, rp := 0, l-1
	for lp < l && rp > -1 {
		res[lp] *= lProduct
		res[rp] *= rProduct

		lProduct *= nums[lp]
		rProduct *= nums[rp]

		lp++
		rp--
	}

	return res
}

func productExceptSelfNaiveButFaster(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}

	product := 1
	var hasDivideByZero bool
	var zeroCnt, zeroIdx int

	res := make([]int, l)

	for idx, num := range nums {
		if num == 0 {
			hasDivideByZero = true
			zeroCnt++

			if zeroCnt > 1 {
				return res
			}
			zeroIdx = idx
			continue
		}

		product *= num
	}

	if hasDivideByZero {
		res[zeroIdx] = product
		return res
	}

	for idx, num := range nums {
		nums[idx] = product / num
	}

	return nums
}
