package linked_list

import "fmt"

type Node struct {
	Value      int
	Prev, Next *Node
}

type LinkedList struct {
	Head, Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *LinkedList) Remove(value int) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		if current.Value == value {
			if current == l.Head {
				l.Head = current.Next
				if l.Head != nil {
					l.Head.Prev = nil
				}
			} else if current == l.Tail {
				l.Tail = current.Prev
				if l.Tail != nil {
					l.Tail.Next = nil
				}
			} else {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			}
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

// func main() {
// 	list := NewLinkedList()
// 	list.Insert(1)
// 	list.Insert(2)
// 	list.Insert(3)
// 	list.Insert(4)
// 	list.Print() // Output: 1 2 3 4
// 	list.Remove(3)
// 	list.Print() // Output: 1 2 4
// }
