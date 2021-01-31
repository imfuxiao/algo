package read

func ProductExceptSelf(nums []int) []int {
	var (
		length = len(nums)
		answer = make([]int, length)
	)

	// 当前数字的左侧, index=0数字没有左侧数字, 所以初始为1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 右侧数字, index=length-1没有右侧数字, 所以它的值为左侧数字乘积, 在上一for循环已完成
	r := nums[length-1] // 保存右侧前数字的乘积
	for i := length - 2; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}

	return answer
}
