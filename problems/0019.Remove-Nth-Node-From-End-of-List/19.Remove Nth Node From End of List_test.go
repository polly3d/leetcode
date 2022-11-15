package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *structs.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{"1,2,3,4,5", args{head: structs.IntsToList([]int{1, 2, 3, 4, 5}), n: 2}, structs.IntsToList([]int{1, 2, 3, 5})},
		{"1", args{head: structs.IntsToList([]int{1}), n: 1}, structs.IntsToList([]int{})},
		{"1,2", args{head: structs.IntsToList([]int{1, 2}), n: 1}, structs.IntsToList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(structs.ListToInt(got), structs.ListToInt(tt.want)) {
				t.Errorf("removeNthFromEnd() = %v, want %v", structs.ListToInt(got), structs.ListToInt(tt.want))
			}
		})
	}
}
