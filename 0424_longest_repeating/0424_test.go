package main

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
