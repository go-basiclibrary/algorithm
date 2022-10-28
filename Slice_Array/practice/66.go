package main

// 66 加一
func plusOne(digits []int) []int {
	stack := make([]int, 0, len(digits))
	for i := 0; i < len(digits); i++ {
		if i == len(digits)-1 {
			stack = append(stack, digits[i]+1)
			break
		}
		stack = append(stack, digits[i])
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == 10 {
			if len(stack)-2 < 0 {
				newSlice := make([]int, 0)
				newSlice = append(newSlice, 1)
				digits[0] = 0
				newSlice = append(newSlice, digits...)
				return newSlice
			}
			stack[len(stack)-2] = stack[len(stack)-2] + 1
			digits[len(stack)-1] = 0
		} else {
			digits[len(stack)-1] = stack[len(stack)-1]
		}
		stack = stack[:len(stack)-1]
	}
	return digits
}

//单for
func plusOne02(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		//从尾部遍历
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func main() {
	plusOne02([]int{0})
}
