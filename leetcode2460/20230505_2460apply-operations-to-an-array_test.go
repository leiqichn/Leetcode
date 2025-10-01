package leetcode2460

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"999", args{[]int{1, 2, 2, 1, 1, 0}}, []int{1, 4, 2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
