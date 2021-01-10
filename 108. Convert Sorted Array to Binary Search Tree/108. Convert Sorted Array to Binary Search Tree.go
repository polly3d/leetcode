package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) <= 0 {
		return nil
	}

	t := &TreeNode{}
	mid := len(nums) / 2
	t.Val = nums[mid]
	t.Left = sortedArrayToBST(nums[:mid])
	t.Right = sortedArrayToBST(nums[mid+1:])
	return t
}
