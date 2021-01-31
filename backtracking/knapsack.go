package backtracking

// weight: 物品重量
// w: 背包可承载重量
func knapsack(weight []int, w int) int {
	n := len(weight) // 物品数量
	return k(weight, w, n, 0, 0)
}

// weight: 物品重量
// maxWeight: 背包最大负重
// n: 物品总数量
// index: 物品索引
// w: 背包当前负重
func k(weight []int, maxWeight, n, index, w int) int {
	if index >= n || w == maxWeight { // w 不可能比 maxWeight大, 因为下面有判断: if weight[index]+w <= maxWeight {
		return w
	}

	r := w

	weightNotPut := k(weight, maxWeight, n, index+1, w) // 不装index+1物品
	if weightNotPut > r {
		r = weightNotPut
	}

	if weight[index]+w <= maxWeight {
		weightInput := k(weight, maxWeight, n, index+1, w+weight[index]) // 装入第index个商品
		if weightInput > r {
			r = weightInput
		}
	}
	return r
}

// weight: 物品重量
// w: 背包可承载重量
func knapsack0(weight []int, w int) int {
	n := len(weight) // n表示物品个数

	status := make([][]bool, n)
	for i := range status {
		status[i] = make([]bool, w+1)
	}

	// 第0个物品需要特殊处理
	status[0][0] = true // 不放第1个物品
	if weight[0] <= w { // 当第1个物品重量小于背包总承重, 则放第一个物品到背包中
		status[0][weight[0]] = true
	}

	// 第1个物品已经特殊处理了, 所以从第2个物品开始
	for i := 1; i < n; i++ { // 动态规划状态转移

		// 不放入第i个物品
		// 把上一行全部重量的状态复制到当前行
		// 因为重量不变表示没有加入当前物品到背包中
		for j := 0; j <= w; j++ {
			if status[i-1][j] {
				status[i][j] = status[i-1][j]
			}
		}

		// 放入第i个物品
		// 取上一行中放入物品的全部重量, 在此的基础上增加当前物品的重量
		// 这里请注意: 终止条件 w - weight[i]表示什么意思呢?
		// 因为只能 j + weight[i] <= w, 即背部装入物品的重量只能小于等于总负重
		// 所以为了让下面 status[i][j+weight[i]] 的二维下标不越界, 提前限制条件 j <= w-weight[i]
		// 写成 j + weight[i] <= w 是不是更好理解呢?
		for j := 0; j <= w-weight[i]; j++ {
			if status[i-1][j] {
				status[i][j+weight[i]] = true
			}
		}
	}

	// 从大到小遍历, 取最后一个最大值
	for i := w; i >= 0; i-- {
		if status[n-1][i] {
			return i
		}
	}
	return 0
}

// weight: 物品重量
// w: 背包总负重
func knapsack2(weight []int, w int) int {
	n := len(weight) // 物品总数量

	status := make([]bool, w+1) // 动态规划中每一层中物品对应的总重量, 下标为重量, 值为是否放入数据

	// 第一个物品特殊处理
	status[0] = true
	if weight[0] <= w {
		status[weight[0]] = true
	}

	for i := 1; i < n; i++ {
		// TODO 因为公用一层status, 所以就不在需要获取上一层的状态到当前层, 即: 不放入当前物品

		// 放入当前物品
		// 必须从大向小进行, 否则会导致重复计算
		// 因为假设status为[true, false, false, false, true, false, false]
		// 从小开始在0处就会+weight[i]的重量, 导致高位false会被改为true
		// 那么当遍历到高位时, 就会进行多于的计算
		for j := w - weight[i]; j >= 0; j-- {
			if status[j] {
				status[j+weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if status[i] {
			return i
		}
	}

	return 0
}

// 背包在不超过最大负重的情况下, 能装入多大价值的物品
// weights: 物品重量
// values: 物品价值
// maxWeight: 背包最大负重
// 返回能装入的最大价值
func knapsack3(weights, values []int, maxWeight int) int {
	n := len(weights)
	return k3(weights, values, maxWeight, n, 0, 0, 0)
}

// n: 物品总数量
// index: 物品的索引
// w: 当前背包负重
// v: 当前背包价值
func k3(weights, values []int, maxWeight, n, index, w, v int) int {
	if index >= n || w == maxWeight {
		return v
	}

	value := v

	// 不放入index+1物品
	valueNotPut := k3(weights, values, maxWeight, n, index+1, w, v)
	if valueNotPut > value {
		value = valueNotPut
	}

	if weights[index]+w <= maxWeight {
		valueInPut := k3(weights, values, maxWeight, n, index+1, w+weights[index], v+values[index])
		if valueInPut > value {
			value = valueInPut
		}
	}

	return value
}

// 背包在不超过最大负重的情况下, 能装入多大价值的物品
// weights: 物品重量
// values: 物品价值
// maxWeight: 背包最大负重
// 返回能装入的最大价值
func knapsack4(weights, values []int, maxWeight int) int {
	n := len(weights)                  // n 表示物品总数量
	status := make([]int, maxWeight+1) // 保存每层的状态, 下标为当前背包重量, 值为背包中物品价值
	for i := range status {
		status[i] = -1
	}

	// 初始化第一层, 即第一个物品
	status[0] = 0
	if weights[0] < maxWeight {
		status[weights[0]] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := maxWeight - weights[i]; j >= 0; j-- {
			if status[j] != -1 {
				value := status[j] + values[i]
				if value > status[j+weights[i]] {
					status[j+weights[i]] = value
				}
			}
		}
	}

	maxValue := 0
	for i := range status {
		if status[i] > maxValue {
			maxValue = status[i]
		}
	}
	return maxValue
}
