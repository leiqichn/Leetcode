package leetcode225

import "testing"

func TestMyStack_Pop(t *testing.T) {
	type fields struct {
		inputQueue  []int
		outputQueue []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"testpop", fields{[]int{1, 2}, []int{}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyStack{
				inputQueue:  tt.fields.inputQueue,
				outputQueue: tt.fields.outputQueue,
			}
			if got := this.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
