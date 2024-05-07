/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode2511

import "testing"

func Test_captureForts(t *testing.T) {
	type args struct {
		forts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"good", args{forts: []int{1, 0, 0, -1, 0, 0, -1, 0, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captureForts(tt.args.forts); got != tt.want {
				t.Errorf("captureForts() = %v, want %v", got, tt.want)
			}
		})
	}
}
