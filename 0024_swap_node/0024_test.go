package main

import (
	. "main/node"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "A",
			args: args{arr: []int{1, 2, 3, 4}},
			want: NewListNode([]int{2, 1, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(NewListNode(tt.args.arr)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
				t.Error(tt.args.arr)
				t.Error(got.Arr())
				t.Error(tt.want.Arr())
			}
		})
	}
}
