package node

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "first",
			args: args{arr: []int{1, 2, 3, 4}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
		},
		{
			name: "empty",
			args: args{arr: []int{}},
			want: nil,
		},
		{
			name: "empty2",
			args: args{arr: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Arr(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want []int
	}{
		{
			name: "simple",
			head: NewListNode([]int{1, 2, 3, 4}),
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.head.Arr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNode.Arr() = %v, want %v", got, tt.want)
			}
		})
	}
}
