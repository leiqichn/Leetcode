package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"good", args{3, 7}, [][]int{{1, 2, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
