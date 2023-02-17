package main

// 北,东,南,西
var lX = [4]int{0, 1, 0, -1}
var lY = [4]int{1, 0, -1, 0}

// 874 模拟行走机器人
func robotSim(commands []int, obstacles [][]int) int {
	//判断障碍物在当前位置是否存在
	obstaclesMap := make(map[[2]int]bool)
	for _, v := range obstacles {
		obstaclesMap[[2]int{v[0], v[1]}] = true
	}
	x, y, dir, distance := 0, 0, 0, 0 // 当前位置,当前方向,欧式距离
	for _, cmd := range commands {
		switch cmd {
		case -2: // 左转
			dir = (dir + 3) % 4
		case -1: //右转
			dir = (dir + 1) % 4
		default: // 前进
			for i := 0; i < cmd; i++ {
				nextX := x + lX[dir]
				nextY := y + lY[dir]
				// 判断是否存在障碍物
				if _, ok := obstaclesMap[[2]int{nextX, nextY}]; ok {
					break
				}
				x = nextX
				y = nextY
				// 计算最大欧式距离
				distance = maxDistance(distance, x*x+y*y)
			}
		}
	}
	return distance
}

func maxDistance(distance int, i int) int {
	if distance > i {
		return distance
	}
	return i
}
