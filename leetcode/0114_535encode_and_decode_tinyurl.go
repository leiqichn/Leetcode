package leetcode

import (
	"strconv"
	"strings"
)

// 通过自增id 的方式来实现encode
type Codec struct {
	dataId map[int]string
	id     int
}

func Constructor() Codec {
	return Codec{map[int]string{}, 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.id++
	this.dataId[this.id] = longUrl
	res := "http://tinyurl.com/" + strconv.Itoa(this.id)
	return res
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	index := strings.Split(shortUrl, "/")
	tmp := index[len(index)-1]
	idx, _ := strconv.Atoi(tmp)
	long := this.dataId[idx]
	return long
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
