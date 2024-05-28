package binary_heap

import "log"

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{data: []int{}}
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Empty() bool {
	return h.Size() == 0
}

func (h *Heap) Insert(value int) {
	h.data = append(h.data, value)
	h.percolateUp(h.Size() - 1)
}

func (h *Heap) Extract() int {
	if h.Empty() {
		log.Fatal("Heap is empty")
	}

	root := h.data[0]
	last := h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	if !h.Empty() {
		h.data[0] = last
		h.percolateDown(0)
	}
	return root
}

func (h *Heap) percolateUp(i int) {
	value := h.data[i]
	for i > 0 {
		parent := (i - 1) / 2
		if value >= h.data[parent] {
			break
		}
		h.data[i] = h.data[parent]
		i = parent
	}
	h.data[i] = value
}

func (h *Heap) percolateDown(i int) {
	value := h.data[i]
	for {
		left := 2*i + 1
		right := 2*i + 2
		min := i
		if left < h.Size() && h.data[left] < h.data[min] {
			min = left
		}
		if right < h.Size() && h.data[right] < h.data[min] {
			min = right
		}
		if min == i {
			break
		}
		h.data[i] = h.data[min]
		i = min
	}
	h.data[i] = value
}

// func main() {
// 	heap := NewHeap()
// 	heap.Insert(5)
// 	heap.Insert(2)
// 	heap.Insert(1)
// 	heap.Insert(3)
// 	heap.Insert(4)
// 	fmt.Println(heap.Extract()) // Output: 1
// 	fmt.Println(heap.Extract()) // Output: 2
// 	fmt.Println(heap.Extract()) // Output: 3
// 	fmt.Println(heap.Extract()) // Output: 4
// 	fmt.Println(heap.Extract()) // Output: 5
// }
