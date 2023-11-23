package two_pointer

func conainterWithTheMostWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var area int
	dist := len(height) - 1
	lp, rp := 0, len(height)-1

	for lp <= rp {
		l, r := height[lp], height[rp]
		tmp := min(l, r) * dist
		if tmp > area {
			area = tmp
		}

		if l > r {
			rp--
		} else {
			lp++
		}
		dist--
	}

	return area
}
