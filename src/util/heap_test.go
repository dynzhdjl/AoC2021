package util

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := MinHeap([]int{48, 24, 7, 299, 3, 2, 1})
	fmt.Println(heap)

	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(PeekHeap(&heap))
	fmt.Println(heap)
}
