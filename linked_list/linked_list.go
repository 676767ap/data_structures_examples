package linked_list

import "fmt"

type Node struct {
	val        int
	prev, next *Node
}

type MyLinkedList struct {
	lenght     int
	head, tail *Node
}

func NewLinkedList() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		tail:   nil,
		lenght: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	if cur == nil || this.lenght-1 < index {
		return -1
	}
	var idx int
	for idx <= index {
		if idx == index {
			return cur.val
		}
		if cur.next == nil {
			break
		}
		cur = cur.next
		idx++
	}
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val:  val,
		next: this.head,
		prev: nil,
	}
	if this.head != nil {
		this.head.prev = newNode
		if this.lenght == 1 {
			this.tail = this.head
		}
	}
	this.head = newNode
	this.lenght++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
		prev: this.tail,
	}
	if this.lenght == 0 {
		this.head = newNode
		this.lenght++
		return
	}
	if this.tail != nil {
		this.tail.next = newNode
	} else {
		this.head.next = newNode
		newNode.prev = this.head
	}
	this.tail = newNode
	this.lenght++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.head
	if (cur == nil && index != 0) || this.lenght < index {
		return
	}
	var idx int
	for idx <= index {
		if idx == index {
			newNode := &Node{
				val:  val,
				next: cur,
			}
			if idx == 0 {
				newNode.prev = nil
				this.head = newNode
				if cur != nil {
					if this.lenght == 1 {
						this.tail = cur
					}
					cur.prev = newNode
				}
				break
			}
			newNode.prev = cur.prev
			cur.prev.next = newNode
			cur.prev = newNode
		} else if index == this.lenght && idx == this.lenght-1 {
			newNode := &Node{
				val:  val,
				next: nil,
				prev: cur,
			}
			this.tail = newNode
			cur.next = newNode
			break
		}
		cur = cur.next
		idx++
	}
	this.lenght++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head
	if cur == nil || this.lenght <= index {
		return
	}
	if this.lenght == 1 && index == 0 {
		this.head = nil
		this.tail = nil
		this.lenght--
		return
	}
	var idx int
	for idx <= index {
		if idx == index {
			if cur == this.tail {
				prev := cur.prev
				prev.next = nil
				this.tail = prev
			} else if cur == this.head {
				if cur.next == nil {
					this.head = nil
					break
				}
				next := cur.next
				cur.next = nil
				next.prev = nil
				this.head = next
			} else {
				prev := cur.prev
				next := cur.next
				prev.next = next
				next.prev = prev
			}
			break
		}
		cur = cur.next
		idx++
	}
	this.lenght--
}

func (l *MyLinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.val)
		current = current.next
	}
	fmt.Println("len: %d", l.lenght)
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
