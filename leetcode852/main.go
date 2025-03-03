package leetcode852

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1 // 左闭右闭
	for left <= right {          // 左闭右闭则left 可以相等
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] { // 判断符合条件的 直接return
			fmt.Println("res:", mid)
			return mid
		}
		if arr[mid] > arr[mid-1] {
			left = mid
		} else {
			right = mid
		}
	}
	return 0
}
