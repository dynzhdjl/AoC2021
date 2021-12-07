package avl

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

func new(key int) *Node {
	return &Node{key, nil, nil, 0}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func balance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) Insert(key int) *Node {
	return insert(t.Root, key)
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return new(key)
	}
	if key > node.Key {
		node.Right = insert(node.Right, key)
	} else if key < node.Key {
		node.Left = insert(node.Left, key)
	} else {
		return node
	}
	node.Height = 1 + max(height(node.Left), height(node.Right))
	//balance := balance(node)
	return node
}
