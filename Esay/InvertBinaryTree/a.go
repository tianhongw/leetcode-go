package a

import "errors"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type queue struct {
	items []interface{}
	size  int
}

func (q *queue) push(item interface{}) {
	q.items = append(q.items, item)
	q.size++
}

func (q *queue) empty() bool {
	return q.size == 0
}

func (q *queue) pop() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("empty queue")
	}

	x := q.items[0]
	q.size--
	q.items = q.items[1:]

	return x, nil
}

func (q *queue) front() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("empty queue")
	}

	return q.items[0], nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var q queue
	q.push(root)

	for !q.empty() {
		fr, err := q.front()
		if err != nil {
			panic(err)
		}
		q.pop()
		fr1 := fr.(*TreeNode)
		l := fr1.Left
		fr1.Left = fr1.Right
		fr1.Right = l

		if fr1.Left != nil {
			q.push(fr1.Left)
		}
		if fr1.Right != nil {
			q.push(fr1.Right)
		}
	}

	return root
}
