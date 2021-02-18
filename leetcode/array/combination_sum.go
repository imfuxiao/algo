package array

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var (
		length      = len(candidates)
		ans         = make([][]int, 0, length)
		current     = make([]int, 0, length)
		combination func(int, int)
	)

	combination = func(target, index int) {
		if index == length {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, current...))
			return
		}

		combination(target, index+1)

		if target-candidates[index] >= 0 {
			current = append(current, candidates[index])
			combination(target-candidates[index], index)
			current = current[:len(current)-1]
		}
	}

	combination(target, 0)

	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	var (
		length      = len(candidates)
		ans         = make([][]int, 0, length)
		current     = make([]int, 0, length)
		combination func([]int, int, int)
	)

	sort.Ints(candidates)

	combination = func(candidates []int, start, target int) {
		if target == 0 {
			ans = append(ans, append([]int{}, current...))
			return
		}

		for i := start; i < length; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if target < candidates[i] {
				break
			}
			current = append(current, candidates[i])
			combination(candidates, i+1, target-candidates[i])
			current = current[:len(current)-1]
		}
	}

	combination(candidates, 0, target)

	return ans
}
