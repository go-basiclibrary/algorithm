package main

import (
	"fmt"
)

// 移动零
func main() { //[0,1,0,3,12]
	moveZeroes04([]int{0, 1, 0, 3, 12})
}

/*
	方法一:
		截取切片,进行append  维护一个指针++,最后补零  不好,要循环过多的slice(或者维护一个slice).不太好实现
	方法二:
		循环nums 不等于零时,与维护指针j交换位置
		判断j是否和i相等,不相等,说明前面遇到过0,则将i置为0
	方法三:
		引入新数组
	方法四:
		双指针
	方法五:
		滚雪球
*/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			//exchange index
			nums[j] = nums[i]

			if i != j { //说明前面遇到过0
				nums[i] = 0
			}
			j++
		}
	}
}

// 方法3
func moveZeroes02(nums []int) {
	newNums := make([]int, 0)
	zeroPoint := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			newNums = append(newNums, nums[i])
		} else {
			zeroPoint++
		}
	}

	for i := 0; i < zeroPoint; i++ {
		newNums = append(newNums, 0)
	}
	nums = newNums
	fmt.Println(nums)
}

// 方法4
func moveZeroes03(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//方法五
func moveZeroes04(nums []int) {
	snowSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowSize++
		} else if snowSize > 0 {
			nums[i], nums[i-snowSize] = 0, nums[i]
		}
	}
	fmt.Println(nums)
}
