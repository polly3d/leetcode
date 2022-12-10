package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *structs.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{"1,2,3,4", args{structs.IntsToList([]int{1, 2, 3, 4})}, structs.IntsToList([]int{2, 1, 4, 3})},
		{"", args{structs.IntsToList([]int{})}, structs.IntsToList([]int{})},
		{"1", args{structs.IntsToList([]int{1})}, structs.IntsToList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", structs.ListToInt(got), structs.ListToInt(tt.want))
			}
		})
	}
}
