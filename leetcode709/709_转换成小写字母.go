package main

import "fmt"

func toLowerCase(s string) string {
	bytes := []byte{}

	sByte := []byte(s)
	for _, b := range sByte {
		if b >= 'A' && b <= 'Z' {
			fmt.Println(b)
			bytes = append(bytes, b+32)
		} else {
			bytes = append(bytes, b)
		}
	}
	return string(bytes)
}
