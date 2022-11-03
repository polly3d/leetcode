package leetcode

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{s: "babad"}, "aba"},
		{"test", args{s: "cbbd"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{s: "babad"}, "bab"},
		{"test", args{s: "cbbd"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
