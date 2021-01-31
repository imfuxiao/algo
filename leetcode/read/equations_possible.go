package read

func equationsPossible(equations []string) bool {
	var (
		equal       byte = '='
		status           = make([]int, 26+1)
		changeValue func(a int, value int)
	)

	changeValue = func(a int, value int) {
		if a == 0 {
			return
		}
		temp := status[a]
		if temp != 0 {
			for i := range status {
				if status[i] == temp {
					status[i] = value
				}

			}
		}
		status[a] = value
	}
	for i := range equations {
		a, b := int(equations[i][0]-'a')+1, int(equations[i][3]-'a')+1
		if equations[i][1] == equal {
			changeValue(a, b)
			changeValue(b, b)
		} else {
			changeValue(a, a)
			changeValue(b, b)
		}
	}
	var ans = true
	for i := range equations {
		a, b := int(equations[i][0]-'a')+1, int(equations[i][3]-'a')+1
		if equations[i][1] != equal {
			ans = ans && (status[a] != status[b])
		} else {
			ans = ans && (status[a] == status[b])
		}
	}
	return ans
}
