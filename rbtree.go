package main

type rbTreeNode struct {
	value int
	//0 red; 1 black
	color  int
	left   *rbTreeNode
	right  *rbTreeNode
	parent *rbTreeNode
}

func (n *rbTreeNode) leftRotate() {
	if n.right == nil {
		return
	}
	p := n.parent
	r := n.right
	rl := r.left
	n.right = rl
	n.parent = r
	if rl != nil {
		rl.parent = n
	}
	if p != nil && p.value > r.value {
		p.left = r
	} else if p != nil && p.value < r.value {
		p.right = r
	}
	r.parent = p
	r.left = n
}

func (n *rbTreeNode) rightRotate() {
	if n.left == nil {
		return
	}
	p := n.parent
	l := n.left
	lr := l.right
	n.left = lr
	n.parent = l
	if lr != nil {
		lr.parent = n
	}
	if p.value > l.value {
		p.left = l
	} else {
		p.right = l
	}
	l.parent = p
	l.right = n
}

type rbTree struct {
	root *rbTreeNode
}

func (t *rbTree) Add(value int) {
	if t.root == nil {
		t.root = createRBTreeNode(value, 1)
		return
	}

	p := t.findNode(value)
	if p.value == value {
		return
	}
	c := createRBTreeNode(value, 0)
	c.parent = p
	pp := p.parent
	if c.value < p.value {
		p.left = c
	} else {
		p.right = c
	}
	for {
		p = c.parent
		if p != nil {
			pp = p.parent
		} else {
			pp = nil
		}
		if p == nil {
			c.color = 1
			break
		}
		if p.color == 1 {
			break
		}
		isPLeft := pp.left != nil && pp.left.value == p.value
		isCLeft := p.left != nil && p.left.value == c.value
		s := pp.left
		if isPLeft {
			s = pp.right
		}

		if s != nil && s.color == 0 {
			p.color = 1
			s.color = 1
			pp.color = 0
			c = pp
		} else {
			if isPLeft {
				if !isCLeft {
					p.leftRotate()
				}
				pp.rightRotate()
				pp.parent.color = 1
				pp.color = 0
			} else {
				if isCLeft {
					p.rightRotate()
				}
				pp.leftRotate()
				pp.parent.color = 1
				pp.color = 0
			}
			break
		}
	}
}

func (t *rbTree) Delete(value int) {
	node := t.findNode(value)
	if node.value == value {
		return
	}
}

func (t *rbTree) Find(value int) *rbTreeNode {
	node := t.findNode(value)
	if node != nil && node.value != value {
		return nil
	}
	return node
}

func (t *rbTree) findNode(value int) *rbTreeNode {
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

func createRBTreeNode(value, color int) *rbTreeNode {
	return &rbTreeNode{
		value: value,
		color: color,
	}
}

func CreateRBTree(value int) *rbTree {
	return &rbTree{
		root: createRBTreeNode(value, 1),
	}
}
