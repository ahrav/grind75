package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], i}
		}
		m[num] = i
	}

	return nil
}
