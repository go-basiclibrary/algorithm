package main

// 22 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var generate func(l, r, n int, s string)
	generate = func(l, r, n int, s string) {
		if l == n && r == n {
			res = append(res, s)
			return
		}
		if l < n {
			generate(l+1, r, n, s+"(")
		}
		if r < l {
			generate(l, r+1, n, s+")")
		}
	}
	generate(0, 0, n, "")
	return res
}
