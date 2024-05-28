package main

import "main/linked_list"

func main() {
	list := linked_list.NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Print() // Output: 1 2 3 4
	list.Remove(3)
	list.Print() // Output: 1 2 4
}
