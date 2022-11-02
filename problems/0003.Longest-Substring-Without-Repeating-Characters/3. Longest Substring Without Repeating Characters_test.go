package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
		{"", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.arg); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
