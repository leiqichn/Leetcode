package leetcode

var (
	m    []string
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	// 将index 和 字符串对应起来
	m = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	backTracking(digits, 1)
	return res
}
func backTracking(digits string, start int) {
	// 终止条件 ，遍历完digits
	if start == len(path) {
		var tmp []byte
		tmp = make([]byte, len(digits))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	// 找到 digit
	digitNum := int(digits[start] - '0')
	// 遍历 digit 对应的map 字符串
	str := m[digitNum]
	for i := 0; i < len(str); i++ { // i 从0开始，因为每个字典都是一个独立的集合，之前的组合是一个集合，所以才从start 开始
		path = append(path, str[i])
		backTracking(digits, start+1)
		path = path[:len(path)-1]
	}
}
