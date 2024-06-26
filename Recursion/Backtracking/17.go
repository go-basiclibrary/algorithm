package main

// 17 电话号码的字母组合

var letterM = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// func1 递归回溯
//func letterCombinations(digits string) []string {
//	var res = make([]string, 0, len(digits)^3)
//
//	var dfs func(i int, s string)
//	dfs = func(index int, s string) {
//		// terminate
//		if len(s) == len(digits) {
//			res = append(res, s)
//			return
//		} else {
//			var digit = string(digits[index])
//			var letter = letterM[digit]
//			for i := 0; i < len(letter); i++ {
//				dfs(index+1, s+string(letter[i]))
//			}
//		}
//	}
//	dfs(0, "")
//	return res
//}
