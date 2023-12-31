package two_pointer

func containerWithTheMostWater(h []int) int {
	if len(h) == 0 {
		return 0
	}

	var area int
	dist := len(h) - 1
	lp, rp := 0, len(h)-1

	for lp <= rp {
		l, r := h[lp], h[rp]
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
