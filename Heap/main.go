package main

type Heap struct {
	root      int
	RightNode *Heap
	LeftNode  *Heap
}

func (heap *Heap) Insert(number int) {
	if number > heap.root {
		saver := heap.root
		heap.root = number
		if heap.LeftNode == nil {
			heap.LeftNode = &Heap{saver, nil, nil}
		} else {
			heap.RightNode = &Heap{saver, nil, nil}
		}
	} else {
		if heap.LeftNode == nil {
			heap.LeftNode = &Heap{number, nil, nil}
		} else {
			heap.RightNode = &Heap{number, nil, nil}
		}
	}
}
func main() {
	heap := Heap{root: 12, RightNode: nil, LeftNode: nil}

	heap.Insert(23)

}
