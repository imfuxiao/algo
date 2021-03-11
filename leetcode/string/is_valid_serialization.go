package string

func isValidSerialization(preorder string) bool {
	var (
		length = len(preorder)
		slot   = 1
	)

	for i := 0; i < length; i++ {
		if slot == 0 { // 当字符串遍历期间, 槽位为空, 则说明树结构不合法
			return false
		}
		switch preorder[i] {
		case ',':
		case '#': // 空节点, 消耗一个槽位
			slot--
		default: // 非空节点, 消耗一个槽位, 并补充两个槽位 slot = slot - 1 + 2 = slot + 1
			// 当节点值大于9时
			for i < length && preorder[i] >= '0' && preorder[i] <= '9' {
				i++
			}
			i--
			slot++
		}
	}

	// 当字符串遍历结束后, 槽位为空, 则说明树结构合法, 否则不合法
	return slot == 0
}
