package main

// 回溯递归
// 22 括号生成
func generateParenthesis02(n int) []string {
	var res []string
	var generate func(n, l, r int, str string)
	generate = func(n, l, r int, str string) {
		// terminate
		if l == n && r == n {
			res = append(res, str)
			return
		}

		// cur logic
		if l < n {
			generate(n, l+1, r, str+"(")
		}

		if r < l {
			generate(n, l, r+1, str+")")
		}
	}
	generate(n, 0, 0, "")
	return res
}
