package leetcode271

import "strings"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	res := ""
	tmpStrs := ""
	lens := string(len(strs))
	tmpStrs = tmpStrs + lens
	for i := 0; i < len(strs); i++ {
		tmpStrs += string(len(strs[i]))
	}
	joinStr := strings.Join(strs, "")
	tmpStrs += joinStr
	tmpByteList := []byte(tmpStrs)
	for i := 0; i < len(tmpByteList); i++ {
		res += string(int(tmpByteList[i]))
	}
	return res
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	lens := byte2int(strs[0])
	end := lens+1
	res := []string{}
	for i:=1;i <end;i++ {
		start := end
		end := start + byte2int(strs[i])
		tmpStr := strs[start,end]

	}
}

func intStr2ByteStr(intStr string) string {
	res := ""
	for i:=0; i<len(intStr);i++ {
		tmpByte := byte(intStr[i])
		res + =1
	}
}

func byte2int(a byte) (b int) {
	return int(a - '0')
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
