package leetcode1700

// 9分钟完成
func countStudents(students []int, sandwiches []int) int {
	// 栈模拟
	// 结束条件
	// 同学中数字都相同，且不等于栈顶元素 [0]

	for !isEnd(students, sandwiches) {
		if sandwiches[0] == students[0] {
			sandwiches = sandwiches[1:]
			students = students[1:]
		} else {
			students = append(students[1:], students[0])
		}
	}

	return len(students)
}

func isEnd(students []int, sandwiches []int) bool {
	for _, val := range students {
		if val == sandwiches[0] {
			return false
		}
	}
	return true
}
