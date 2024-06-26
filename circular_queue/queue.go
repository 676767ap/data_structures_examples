package circularqueue

type MyCircularQueue struct {
	queue  []int
	head   int
	tail   int
	lenght int
	size   int
}

func Constructor(k int) MyCircularQueue {
	myQueue := MyCircularQueue{
		queue:  make([]int, k),
		head:   -1,
		tail:   -1,
		lenght: k,
		size:   0,
	}
	return myQueue
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.tail == this.lenght-1 {
		this.tail = 0
		this.queue[this.tail] = value
		this.size++
		return true
	}
	this.queue[this.tail+1] = value
	this.tail++
	if this.head == -1 {
		this.head = 0
	}
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.queue[this.head] = -1
	this.size--
	if this.IsEmpty() {
		this.head = -1
		this.tail = -1
		return true
	}
	if this.head == this.lenght-1 {
		this.head = 0
	} else {
		this.head++
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.lenght
}
