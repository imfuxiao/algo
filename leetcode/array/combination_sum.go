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

func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans := make([][]int, 0, 10)
	current := make([]int, 0, k)

	var combination func(int, int)

	combination = func(target int, index int) {
		if target == 0 && len(current) == k {
			ans = append(ans, append([]int{}, current...))
			return
		}
		for i := index; i < 9; i++ {
			if nums[i]-target > 0 {
				break
			}
			if len(current) >= k {
				return
			}
			current = append(current, nums[i])
			combination(target-nums[i], i+1)
			current = current[:len(current)-1]
		}
	}

	combination(n, 0)

	return ans
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
