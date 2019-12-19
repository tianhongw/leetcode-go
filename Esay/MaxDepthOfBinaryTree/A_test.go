package a

import "testing"

func TestMaxDepth(t *testing.T) {
	root := new(TreeNode)
	n1 := TreeNode{
		val: 9,
	}
	n2 := TreeNode{
		val: 20,
	}
	n3 := TreeNode{
		val: 15,
	}
	n4 := TreeNode{
		val: 7,
	}

	root.left = &n1
	root.right = &n2
	n2.left = &n3
	n2.right = &n4

	if maxDepth(root) != 3 {
		t.Fail()
	}
}
