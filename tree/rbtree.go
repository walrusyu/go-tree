package tree

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

func (n *rbTreeNode) nextNode() *rbTreeNode {
	if n == nil || n.right == nil {
		return nil
	}
	c := n.right
	for c.left != nil {
		c = c.left
	}
	return c
}

func (n *rbTreeNode) up() {
	var p, pp *rbTreeNode
	for {
		p = n.parent
		if p != nil {
			pp = p.parent
		} else {
			pp = nil
		}
		if p == nil {
			n.color = 1
			break
		}
		if p.color == 1 {
			break
		}
		isPLeft := pp.left != nil && pp.left.value == p.value
		isCLeft := p.left != nil && p.left.value == n.value
		s := pp.left
		if isPLeft {
			s = pp.right
		}

		if s != nil && s.color == 0 {
			p.color = 1
			s.color = 1
			pp.color = 0
			n = pp
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

func (n *rbTreeNode) delete() {
	if n.left != nil && n.right != nil {
		next := n.nextNode()
		n.value, next.value = next.value, n.value
		next.delete()
	} else if n.left != nil || n.right != nil {
		// n has one child, n must be black, child must be red
		c := n.left
		if c == nil {
			c = n.right
		}
		n.value, c.value = c.value, n.value
		c.delete()
	} else {
		// n has no child
		n.adjust()
		n.drop()
	}
}

func (n *rbTreeNode) adjust() {
	for {
		p := n.parent
		if p == nil || n.color == 0 {
			// n is root or n is red
			n.color = 1
			break
		} else {
			// n is not root and n is black, s must not be nil
			isLeft := p.left != nil && p.left.value == n.value
			s := p.left
			sn, sf := s.left, s.right
			if isLeft {
				s = p.right
				sn, sf = s.right, s.left
			}

			if s.color == 0 {
				// s is red and s has two black child
				s.color = 1
				p.color = 0
				if isLeft {
					p.leftRotate()
				} else {
					p.rightRotate()
				}
			} else {
				//s is black
				if sn.isBlack() && sf.isBlack() {
					s.color = 0
				} else if sf.isBlack() && !sn.isBlack() {
					s.color, sn.color = sn.color, s.color
					if isLeft {
						s.rightRotate()
					} else {
						s.leftRotate()
					}
				} else if !sf.isBlack() && sn.isBlack() {
					s.color = p.color
					p.color = 1
					sf.color = 1
					if isLeft {
						p.leftRotate()
					} else {
						p.rightRotate()
					}
				}
			}
			n = n.parent
		}
	}
}

func (n *rbTreeNode) isBlack() bool {
	return n == nil || n.color == 1
}

func (n *rbTreeNode) drop() {
	if n.parent != nil {
		p := n.parent
		isLeft := p.left != nil && p.left.value == n.value
		if isLeft {
			p.left = nil
		} else {
			p.right = nil
		}
	}
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
	if c.value < p.value {
		p.left = c
	} else {
		p.right = c
	}
	c.up()
}

func (t *rbTree) Delete(value int) {
	node := t.findNode(value)
	if node.value != value {
		return
	}
	node.delete()
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
