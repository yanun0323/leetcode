package main

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "aaba",
			args: args{
				board: [][]byte{{'a', 'b'}, {'a', 'a'}},
				words: []string{"aaba"},
			},
			want: []string{"aaba"},
		},
		{
			name: "looping",
			args: args{
				board: [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}},
				words: []string{"eaafgdcba"},
			},
			want: []string{"eaafgdcba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
