package _0251006

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 不打破原有规则的情况下种植n 朵花
	// 0 ， 和1 不相邻，

	// 贪心， 从头遍历不相邻的0，置为1

	if n == 0 { // 注意判断特殊值
		return true
	}
	count := n
	for i, _ := range flowerbed {
		if isValid(flowerbed, i) {
			flowerbed[i] = 1
			count--
			if count <= 0 {
				return true
			}
		}
	}
	return false
}

//
//func isVaild(flowerbed []int, put int) bool {
//	n := len(flowerbed) - 1
//	front := put - 1
//	back := put + 1
//
//	if flowerbed[put] != 0 {
//		return false
//	}
//	flag := true
//	if front >= 0 && flowerbed[front] != 1 {
//		flag = true
//	} else if front < 0 {
//		flag = true
//	} else if front >= 0 && flowerbed[front] == 1 {
//		flag = false
//
//	}
//
//	if back <= n && flowerbed[back] != 1 {
//		flag = flag && true
//	} else if back > n {
//		flag = flag && true
//	} else if back <= n && flowerbed[back] == 1 {
//		flag = false
//	}
//
//	return flag
//
//}

// 优化，前面只写false的情况， 则不用合并

// 判断位置i是否可以种花
func isValid(flowerbed []int, i int) bool {
	// 当前位置必须为空
	if flowerbed[i] != 0 {
		return false
	}

	// 检查左边位置
	if i > 0 && flowerbed[i-1] == 1 {
		return false
	}

	// 检查右边位置
	if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
		return false
	}

	return true
}
