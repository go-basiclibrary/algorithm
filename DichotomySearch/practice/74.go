package main

import "fmt"

func main() {
	res := searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 13)
	fmt.Println(res)
}

// 74 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// 二分寻找所在子切片
	var index int
	l, r := 0, len(matrix)-1
	for l <= r {
		mid := (l + r) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			index = mid
			break
		}
		if target < matrix[mid][0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 找到index,进行二分
	iL, iR := 0, len(matrix[index])-1
	for iL <= iR {
		mid := (iL + iR) / 2
		if matrix[index][mid] == target {
			return true
		}
		if matrix[index][mid] < target {
			iL = mid + 1
		} else {
			iR = mid - 1
		}
	}
	return false
}
