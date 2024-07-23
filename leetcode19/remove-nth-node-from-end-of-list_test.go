/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/26 下午11:42
 */

package leetcode19

import (
	. "lcutils/listnode"
	"reflect"
	"testing"
)

func Test_findFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"test1", args{&ListNode{1,&ListNode{2, &ListNode{3,&ListNode{4,&ListNode{5,nil}}}}}},&ListNode{1}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
