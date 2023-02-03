package main

// 22 括号生成(递归回溯)
func generateParenthesis02(n int) []string {
	var res []string
	var generate func(n, l, r int, str string)
	generate = func(n, l, r int, str string) {
		// terminate
		if l == n && r == n {
			res = append(res, str)
			return
		}

		// current logic
		if l < n { // 层层回溯,进行递归
			// next floor
			generate(n, l+1, r, str+"(")
		}
		if r < l { // 限制r数永远小于l即可
			// next floor
			generate(n, l, r+1, str+")")
		}
	}

	// call
	generate(n, 0, 0, "")

	return res
}
