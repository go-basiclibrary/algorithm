package main

import "fmt"

func main() {
	res := uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	fmt.Println(res)
}

// 首先从第一行开始遍历,f[0]已经进行初始化,想到到达obstacleGrid[0][j]只能够从左往右走,
// 上面没有路可以走,java整型数组初始化的时候默认值是零,也就是说f[j]的默认值都是零,
// 经过f[j] += f[j - 1] (此时i=0也就是在第一行)的第一次循环完成后第一行的所有值就已经确定了,
// 没有障碍物就是全为1,有障碍物就是为1或0。现在f[i]数组中保存的值就是从起点出发到第一行第i个网格有多少条不同的路径。
// 之后继续循环，当i=1也就是遍历到第二行，
// （每一行第一个网格的值如果有障碍就将其置0，如果无障碍的话值与上一行的f[j]保持一致，不用变化），
// 第二个网格的值计算方式为上方的网格值加上左边网格的值，左边网格值已经确定，为f[j-1]，
// 而上方网格的值就是上一轮循环中的f[j]，两个值相加得到第二行第二个网格的值并将f[j]中的数据覆盖，
// 依次循环将f[j]中的数据全部替换，此时f[j]中就保存了从起点出发到第二行第j个网格有多少条不同的路径。
// 继续循环直到完成，最后的结果就是从起点到终点有多少条不同的路径。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 { // 遇到障碍
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}

	return f[len(f)-1]
}
