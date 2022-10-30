package leetcode

import (
	"leetcode/structs"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}}, []int{7, 0, 8}},
		{"test", args{l1: []int{2, 4, 8}, l2: []int{5, 6, 4}}, []int{7, 0, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(structs.IntsToList(tt.args.l1), structs.IntsToList(tt.args.l2)); !reflect.DeepEqual(structs.ListToInt(got), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", structs.ListToInt(got), tt.want)
			}
		})
	}
}
