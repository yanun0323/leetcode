package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "aghalfgasg",
			},
			want: false,
		},
		{
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			args: args{
				s: " ",
			},
			want: true,
		},
		{
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			args: args{
				s: "a.",
			},
			want: true,
		},
		{
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
