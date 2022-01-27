package tree

import "testing"

func TestAvlTree_Add(t *testing.T) {
	tree := CreateAVLTree()
	tree.Add(3)
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(10)
	tree.Add(9)
	tree.Add(8)
	if tree.root == nil || tree.root.value != 4 {
		t.Errorf("incorrect node root")
	} else {
		if tree.root.left == nil || tree.root.left.value != 2 {
			t.Errorf("incorrect node root->left")
		} else if tree.root.right == nil || tree.root.right.value != 7 {
			t.Errorf("incorrect node root->right")
		} else {
			if tree.root.left.left == nil || tree.root.left.left.value != 1 {
				t.Errorf("incorrect node root->left>left")
			} else if tree.root.left.right == nil || tree.root.left.right.value != 3 {
				t.Errorf("incorrect node root->left>right")
			} else if tree.root.right.left == nil || tree.root.right.left.value != 6 {
				t.Errorf("incorrect node root->right>left")
			} else if tree.root.right.right == nil || tree.root.right.right.value != 9 {
				t.Errorf("incorrect node root->right>right")
			} else {
				if tree.root.right.left.left == nil || tree.root.right.left.left.value != 5 {
					t.Errorf("incorrect node root->right>left>left")
				} else if tree.root.right.right.left == nil || tree.root.right.right.left.value != 8 {
					t.Errorf("incorrect node root->right>right>left")
				} else if tree.root.right.right.right == nil || tree.root.right.right.right.value != 10 {
					t.Errorf("incorrect node root->right>right>right")
				}
			}
		}
	}
}

func TestAvlTree_Delete(t *testing.T) {
	tree := CreateAVLTree()
	tree.Add(3)
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(10)
	tree.Add(9)
	tree.Add(8)
	tree.Delete(8)
	tree.Delete(5)
	tree.Delete(6)
	if tree.root.right.left == nil || tree.root.right.left.value != 7 || tree.root.right.right == nil || tree.root.right.right.value != 10 {
		t.Errorf("incorrect node root->right>left>left")
	}
}
