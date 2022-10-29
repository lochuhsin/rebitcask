package models

type bstnode struct {
	key   string
	val   []byte
	left  *bstnode
	right *bstnode
}
type BinarySearchTree struct {
	root *bstnode
	size int
}

func (bst *BinarySearchTree) Init() {
	bst.size = 0
}

func (bst *BinarySearchTree) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTree) GetAll() []KVPair {
	kvPair := make([]KVPair, 0, bst.size)
	bst.inorder(bst.root, &kvPair)
	return kvPair
}

func (bst *BinarySearchTree) inorder(root *bstnode, kvPair *[]KVPair) {
	if root == nil {
		return
	}

	if root.left != nil {
		bst.inorder(root.left, kvPair)
	}

	*kvPair = append(*kvPair, KVPair{
		Key: root.key,
		Val: root.val,
	})

	if root.right != nil {
		bst.inorder(root.right, kvPair)
	}

}

func (bst *BinarySearchTree) Set(key string, val []byte) {
	bst.root = bst.set(bst.root, &key, &val)
}

func (bst *BinarySearchTree) Get(key string) (val []byte, status bool) {
	if res := bst.get(bst.root, &key); res != nil {
		return *res, true
	}
	return []byte(""), false
}

func (bst *BinarySearchTree) set(root *bstnode, key *string, val *[]byte) *bstnode {
	if root == nil {
		bst.size++
		return &bstnode{
			key: *key,
			val: *val,
		}
	}

	if root.key == *key {
		root.val = *val
		return root
	}

	if node := *root; node.key != *key {
		if *key < node.key {
			root.left = bst.set(node.left, key, val)
		} else {
			root.right = bst.set(node.right, key, val)
		}
		return root
	}
	return nil
}

func (bst *BinarySearchTree) get(root *bstnode, key *string) (val *[]byte) {
	if root == nil {
		return nil
	}

	if root.key == *key {
		return &root.val
	}

	if *key < root.key {
		return bst.get(root.left, key)
	} else {
		return bst.get(root.right, key)
	}
}
