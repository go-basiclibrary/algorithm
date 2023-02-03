package main

// 443
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	m := make(map[string]struct{})
	for _, v := range bank {
		m[v] = struct{}{}
	}
	if _, ok := m[endGene]; !ok {
		return -1
	}

	q := []string{startGene}

	//DFS思想解决问题,将要变的string想成一棵树
	for step := 0; q != nil; step++ {
		qStr := q[0]
		q = nil
		for i, x := range qStr {
			// x != y的情况下,进行替换
			for _, v := range "ACGT" {
				if x != v {
					newStr := qStr[:i] + string(v) + qStr[i+1:]
					if _, ok := m[newStr]; ok {
						if newStr == endGene {
							return step + 1
						}
						delete(m, newStr)
						q = append(q, newStr)
					}
				}
			}
		}
	}
	return -1
}
