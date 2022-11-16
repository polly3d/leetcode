package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *structs.ListNode
		list2 *structs.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{"[1,2,4],[1,3,4]", args{list1: structs.IntsToList([]int{1, 2, 4}), list2: structs.IntsToList([]int{1, 3, 4})}, structs.IntsToList([]int{1, 1, 2, 3, 4, 4})},
		{"[], []", args{list1: structs.IntsToList([]int{}), list2: structs.IntsToList([]int{})}, structs.IntsToList([]int{})},
		{"[], [0]", args{list1: structs.IntsToList([]int{}), list2: structs.IntsToList([]int{0})}, structs.IntsToList([]int{0})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(structs.ListToInt(got), structs.ListToInt(tt.want)) {
				t.Errorf("mergeTwoLists() = %v, want %v", structs.ListToInt(got), structs.ListToInt(tt.want))
			}
		})
	}
}
