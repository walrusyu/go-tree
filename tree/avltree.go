package tree

import "github.com/walrusyu/go-tree/utils"

type avlTreeNode struct {
	value  int
	left   *avlTreeNode
	right  *avlTreeNode
	height int
}

func (n *avlTreeNode) getHeight() int {
	if n == nil {
		return 0
	} else {
		return n.height
	}
}

func (n *avlTreeNode) refreshHeight() {
	if n == nil {
		return
	}
	left := n.left.getHeight()
	right := n.right.getHeight()
	n.height = utils.Max(left, right) + 1
}

func (n *avlTreeNode) leftRotate() *avlTreeNode {
	if n == nil || n.right == nil {
		return nil
	}
	r := n.right
	n.right = r.left
	r.left = n
	n.refreshHeight()
	r.refreshHeight()
	return r
}

func (n *avlTreeNode) rightRotate() *avlTreeNode {
	if n == nil || n.left == nil {
		return nil
	}
	l := n.left
	n.left = l.right
	l.right = n
	n.refreshHeight()
	l.refreshHeight()
	return l
}

func (n *avlTreeNode) leftRightRotate() *avlTreeNode {
	n.left = n.left.leftRotate()
	return n.rightRotate()
}

func (n *avlTreeNode) rightLeftRotate() *avlTreeNode {
	n.right = n.right.rightRotate()
	return n.leftRotate()
}

func (n *avlTreeNode) insert(value int) *avlTreeNode {
	if n == nil {
		return nil
	}
	if n.value == value {
		return n
	}
	if value < n.value {
		if n.left == nil {
			n.left = createAVLTreeNode(value, 1)
		} else {
			n.left = n.left.insert(value)
			if n.left.getHeight()-n.right.getHeight() == 2 {
				if value < n.left.value {
					//LL
					n = n.rightRotate()
				} else {
					//LR
					n = n.leftRightRotate()
				}
			}
		}
	} else {
		if n.right == nil {
			n.right = createAVLTreeNode(value, 1)
		} else {
			n.right = n.right.insert(value)
			if n.right.getHeight()-n.left.getHeight() == 2 {
				if value < n.right.value {
					//RL
					n = n.rightLeftRotate()
				} else {
					//RR
					n = n.leftRotate()
				}
			}
		}
	}
	n.refreshHeight()
	return n
}

type avlTree struct {
	root *avlTreeNode
}

func (t *avlTree) Add(value int) {
	if t.root == nil {
		t.root = createAVLTreeNode(value, 1)
		return
	}

	t.root = t.root.insert(value)
}

func (t *avlTree) findNode(value int) *avlTreeNode {
	root := t.root
	for root != nil {
		if root.value == value {
			return root
		} else if root.value < value {
			if root.right == nil {
				return root
			} else {
				root = root.right
			}
		} else {
			if root.left == nil {
				return root
			} else {
				root = root.left
			}
		}
	}
	return root
}

func createAVLTreeNode(value, height int) *avlTreeNode {
	return &avlTreeNode{
		value:  value,
		height: height,
	}
}

func CreateAVLTree() *avlTree {
	return &avlTree{}
}
