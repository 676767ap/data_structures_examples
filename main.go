package main

import (
	circularqueue "main/circular_queue"
)

func main() {
	queue := circularqueue.Constructor(2)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.DeQueue()
	queue.EnQueue(3)
	queue.DeQueue()
	queue.EnQueue(3)
	queue.DeQueue()
	queue.EnQueue(3)
	queue.DeQueue()
}
