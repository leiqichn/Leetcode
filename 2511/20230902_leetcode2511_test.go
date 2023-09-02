package _511

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
		// TODO: Add test cases.
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
