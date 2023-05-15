package leetcode

import "strings"

func defangIPaddr(address string) string {
	adressByte := []byte(address) // string 转为byte数组
	res := []byte{}
	for _, str := range adressByte {
		if str == '.' {
			res = append(res, '[')
			res = append(res, '.')
			res = append(res, ']')

		} else {
			res = append(res, str)
		}
	}
	return strings.ReplaceAll(address, ".", "[.]")
	//return string(res)
}
