package bst

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}

	t.insertNode(t.Root, value)
}

func (t *Tree) insertNode(node *Node, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			t.insertNode(node.Left, value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			t.insertNode(node.Right, value)
		}
	}
}

func (t *Tree) Search(value int) bool {
	return t.searchNode(t.Root, value)
}

func (t *Tree) searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return t.searchNode(node.Left, value)
	} else {
		return t.searchNode(node.Right, value)
	}
}

// func main() {
// 	tree := Tree{}

// 	tree.Insert(5)
// 	tree.Insert(3)
// 	tree.Insert(7)
// 	tree.Insert(1)
// 	tree.Insert(4)
// 	tree.Insert(6)
// 	tree.Insert(9)

// 	fmt.Println(tree.Search(7))  // Output: true
// 	fmt.Println(tree.Search(10)) // Output: false
// }
