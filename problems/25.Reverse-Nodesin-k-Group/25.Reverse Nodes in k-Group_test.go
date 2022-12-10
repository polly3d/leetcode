package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *structs.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{"[1,2,3,4,5]", args{head: structs.IntsToList([]int{1, 2, 3, 4, 5}), k: 2}, structs.IntsToList([]int{2, 1, 4, 3, 5})},
		{"[1,2,3,4,5]", args{head: structs.IntsToList([]int{1, 2, 3, 4, 5}), k: 3}, structs.IntsToList([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			ga := structs.ListToInt(got)
			wa := structs.ListToInt(tt.want)
			if !reflect.DeepEqual(ga, wa) {
				t.Errorf("reverseKGroup() = %v, want %v", ga, wa)
			}
		})
	}
}

func Test_reverseKGroup2(t *testing.T) {
	type args struct {
		head *structs.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *structs.ListNode
	}{
		{"[1,2,3,4,5]", args{head: structs.IntsToList([]int{1, 2, 3, 4, 5}), k: 2}, structs.IntsToList([]int{2, 1, 4, 3, 5})},
		{"[1,2,3,4,5]", args{head: structs.IntsToList([]int{1, 2, 3, 4, 5}), k: 3}, structs.IntsToList([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup2(tt.args.head, tt.args.k)
			ga := structs.ListToInt(got)
			wa := structs.ListToInt(tt.want)
			if !reflect.DeepEqual(ga, wa) {
				t.Errorf("reverseKGroup() = %v, want %v", ga, wa)
			}
		})
	}
}
