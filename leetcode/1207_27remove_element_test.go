package leetcode

import "testing"

func Test_removeElement2(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "good", args: args{nums: []int{1, 3, 4, 2}, val: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement2(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}
