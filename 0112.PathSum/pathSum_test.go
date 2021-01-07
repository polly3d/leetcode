package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	arr := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}
	n := &TreeNode{Val: arr[0]}
	for i := 1; i < len(arr)-2; i += 2 {
		if arr[i] != -1 {
			nn := &TreeNode{Val: arr[i]}
		}
	}
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{root: n, sum: 22}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
