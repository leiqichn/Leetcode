/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/10 下午7:49
 */

package leetcode230

import (
	"lcutils/nilint"
	"lcutils/treenode"
	"testing"
)
import . "lcutils/treenode"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "good", args: args{treenode.FromSlice([]nilint.NilInt{nilint.NewInt(3), nilint.NewInt(1), nilint.NewInt(4), nilint.NewNil(), nilint.NewInt(2)}), 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
