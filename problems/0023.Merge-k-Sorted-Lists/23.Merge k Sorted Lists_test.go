package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*structs.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{
			"[[1,4,5],[1,3,4],[2,6]]",
			args{
				lists: []*structs.ListNode{
					structs.IntsToList([]int{1, 4, 5}),
					structs.IntsToList([]int{1, 3, 4}),
					structs.IntsToList([]int{2, 6}),
				},
			},
			structs.IntsToList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			"[]",
			args{
				lists: []*structs.ListNode{},
			},
			structs.IntsToList([]int{}),
		},
		{
			"[[]]",
			args{
				lists: []*structs.ListNode{
					structs.IntsToList([]int{}),
				},
			},
			structs.IntsToList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(
				structs.ListToInt(got),
				structs.ListToInt(tt.want),
			) {
				t.Errorf("mergeKLists() = %v, want %v", structs.ListToInt(got), structs.ListToInt(tt.want))
			}
		})
	}
}
