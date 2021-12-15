package heap

func MinHeap(arr []int) []int {
	heap := []int{}
	var insert func(key, parent int)
	insert = func(key, parent int) {
		if parent == 0 {
			return
		}
		if heap[parent-1] > heap[key-1] {
			heap[parent-1], heap[key-1] = heap[key-1], heap[parent-1]
			insert(parent, parent/2)
		}
	}
	for i := range arr {
		heap = append(heap, arr[i])
		insert(len(heap), len(heap)/2)
	}
	return heap
}

func PeekHeap(heap *[]int) int {
	if len(*heap) == 0 {
		panic("this should not happened")
	}
	h := (*heap)[0]
	(*heap)[0], (*heap)[len((*heap))-1] = (*heap)[len((*heap))-1], (*heap)[0]
	*heap = (*heap)[:len((*heap))-1]

	var delete func(int)
	delete = func(key int) {
		left := key * 2
		right := key*2 + 1
		if left > len(*heap) {
			return
		}
		getSmallestChild := func() int {
			if right > len(*heap) {
				return left
			}
			if (*heap)[left-1] <= (*heap)[right-1] {
				return left
			}
			return right
		}
		child := getSmallestChild()
		if (*heap)[child-1] < (*heap)[key-1] {
			(*heap)[child-1], (*heap)[key-1] = (*heap)[key-1], (*heap)[child-1]
			delete(child)
		}
	}
	delete(1)
	return h
}
