package main

// 22 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var generate func(l, r, n int, s string)
	generate = func(l, r, n int, s string) {
		// 退出条件
		if l == n && r == n {
			res = append(res, s)
			return
		}
		// 执行逻辑
		if l < n {
			// 进入下一层
			generate(l+1, r, n, s+"(")
		}
		if r < l {
			//进入下一层
			generate(l, r+1, n, s+")")
		}
	}
	generate(0, 0, n, "")
	return res
}
