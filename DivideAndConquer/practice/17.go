package main

func main() {
}

var res []string

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 17 电话号码的字母组合
func letterCombinations(digits string) []string {
	// 每回合初始化res
	res = []string{}
	if len(digits) == 0 {
		return res
	}
	backtrack(digits, 0, "")
	return res
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		letters := phoneMap[string(digit)]
		letterLen := len(letters)
		for i := 0; i < letterLen; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
